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
	pm2Config    *dto.EconsystemConfig
}

func NewPm2Manager(prompter ports.Prompter, sshConnector ports.SSHConnector) ports.Pm2Manager {
	return &Pm2Manager{
		prompter:     prompter,
		sshConnector: sshConnector,
	}
}

func (p *Pm2Manager) RetrievePm2Config(ecosystemConfigPath string) (*dto.EconsystemConfig, error) {

	if p.pm2Config != nil {
		return p.pm2Config, nil
	}

	if _, err := os.Stat(ecosystemConfigPath); os.IsNotExist(err) {
		return nil, err
	}

	tmpJsonPath := "/tmp/ecosystem.json"

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
	`, ecosystemConfigPath, tmpJsonPath)

	// missing tests
	cmd := exec.Command("node", "-e", script)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	if err := utils.ReadJSON(tmpJsonPath, &p.pm2Config); err != nil {
		return nil, err
	}

	return p.pm2Config, nil
}

func (p *Pm2Manager) GetAvailableEnvs() ([]string, error) {
	pm2Config, err := p.RetrievePm2Config(filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE))
	if err != nil {
		return []string{}, err
	}

	if len(pm2Config.Deploy) == 0 {
		return []string{}, err
	}

	var options []string
	for env := range pm2Config.Deploy {
		options = append(options, env)
	}
	return options, nil
}

func (p *Pm2Manager) GetAvailableHosts(env string) ([]string, error) {
	if p.pm2Config == nil {
		return nil, fmt.Errorf("PM2 configuration not loaded")
	}
	return p.pm2Config.ListHostsInEnv(env), nil
}

func (p *Pm2Manager) GetSSHConfig(env, host string) (*dto.SSHConfig, error) {
	pm2Config, err := p.RetrievePm2Config(filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE))
	if err != nil {
		return nil, err
	}

	envConfig, exists := pm2Config.Deploy[env]
	if !exists {
		return nil, err
	}

	return &dto.SSHConfig{
		User: envConfig.User,
		Host: host,
		Port: 22,
	}, nil
}
