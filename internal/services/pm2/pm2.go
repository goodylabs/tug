package pm2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

type Pm2Manager struct {
	prompter     ports.Prompter
	sshConnector ports.SSHConnector
}

func NewPm2Manager(prompter ports.Prompter, sshConnector ports.SSHConnector) *Pm2Manager {
	return &Pm2Manager{
		prompter:     prompter,
		sshConnector: sshConnector,
	}
}

func (p *Pm2Manager) LoadPm2Config(ecosystemConfigPath string, pm2Config *dto.EconsystemConfig) error {

	if _, err := os.Stat(ecosystemConfigPath); os.IsNotExist(err) {
		return fmt.Errorf("Can not load config from file(probably doesn't exist): %s", ecosystemConfigPath)
	}

	tmpPath := "/tmp/ecosystem.json"

	script := fmt.Sprintf(`
		const fs = require("fs");
		const config = require("%s");

		if (config.deploy) {
			for (const key in config.deploy) {
				const deployEntry = config.deploy[key];
				if (typeof deployEntry.host === "string") {
					deployEntry.host = [deployEntry.host];
				}
			}
		}

		fs.writeFileSync("%s", JSON.stringify(config, null, 2));
	`, ecosystemConfigPath, tmpPath)

	// missing tests
	cmd := exec.Command("node", "-e", script)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("running node script to load pm2 config: %w", err)
	}

	if err := utils.ReadJSON(tmpPath, pm2Config); err != nil {
		return fmt.Errorf("cannot read json file %s error: %w", tmpPath, err)
	}

	return nil
}

func (p *Pm2Manager) GetAvailableEnvs() ([]string, error) {
	var pm2Config dto.EconsystemConfig

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	if err := p.LoadPm2Config(ecosystemConfigPath, &pm2Config); err != nil {
		return []string{}, fmt.Errorf("error loading PM2 config: %w", err)
	}

	if len(pm2Config.Deploy) == 0 {
		return []string{}, fmt.Errorf("no environments found in PM2 config")
	}

	var options []string
	for env := range pm2Config.Deploy {
		options = append(options, env)
	}
	return options, nil
}

func (p *Pm2Manager) selectHost(pm2Config *dto.EconsystemConfig, selectedEnv string) (string, error) {
	hosts := pm2Config.Deploy[selectedEnv].Host
	if len(hosts) == 0 {
		return "", fmt.Errorf("no hosts found for selected environment %s", selectedEnv)
	}

	if len(hosts) == 1 {
		return hosts[0], nil
	}

	return p.prompter.ChooseFromList(hosts, "Select host for environment "+selectedEnv)
}

func (p *Pm2Manager) GetSSHConfig(selectedEnv string) (*dto.SSHConfig, error) {

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	if err := p.LoadPm2Config(ecosystemConfigPath, &pm2Config); err != nil {
		return nil, fmt.Errorf("error loading PM2 config: %w", err)
	}

	envConfig, exists := pm2Config.Deploy[selectedEnv]
	if !exists {
		return nil, fmt.Errorf("environment '%s' not found in loaded PM2 config", selectedEnv)
	}

	host, err := p.selectHost(&pm2Config, selectedEnv)
	if err != nil {
		return nil, err
	}

	return &dto.SSHConfig{
		User: envConfig.User,
		Host: host,
		Port: 22,
	}, nil
}
