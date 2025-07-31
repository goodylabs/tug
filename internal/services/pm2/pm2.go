package pm2

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

var tmpJsonPath = "/tmp/ecosystem.json"

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

func (p *Pm2Manager) GetPm2ConfigPath() (string, error) {
	for _, name := range []string{"ecosystem.config.js", "ecosystem.config.cjs"} {
		path := filepath.Join(config.BASE_DIR, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("ecosystem config file not found in %s", config.BASE_DIR)
}

func (p *Pm2Manager) ConvertJsFileToJson(pm2ConfigPath string) error {
	script := p.BuildNodeScript(pm2ConfigPath)
	cmd := exec.Command("node", "-e", script)
	return cmd.Run()
}

func (p *Pm2Manager) BuildNodeScript(path string) string {
	if strings.HasSuffix(path, ".js") {
		return fmt.Sprintf(`
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
		`, path, tmpJsonPath)
	}
	return fmt.Sprintf(`
		import config from "%s";

		const { default: ecosystemConfig } = await import("%s");

		if (ecosystemConfig.deploy) {
			for (const key in ecosystemConfig.deploy) {
				const deployEntry = ecosystemConfig.deploy[key];
				if (typeof deployEntry.host === "string") {
					deployEntry.host = [deployEntry.host];
				}
			}
		}

		const fs = await import("fs/promises");
		await fs.writeFile("%s", JSON.stringify(ecosystemConfig, null, 2));
	`, path, path, tmpJsonPath)
}

func (p *Pm2Manager) RetrievePm2Config() (*dto.EconsystemConfig, error) {
	if p.pm2Config != nil {
		return p.pm2Config, nil
	}

	path, err := p.GetPm2ConfigPath()
	if err != nil {
		return nil, err
	}

	if err := p.ConvertJsFileToJson(path); err != nil {
		return nil, err
	}

	var config dto.EconsystemConfig
	if err := utils.ReadJSON(tmpJsonPath, &config); err != nil {
		return nil, err
	}

	p.pm2Config = &config
	return p.pm2Config, nil
}
func (p *Pm2Manager) GetAvailableEnvs() ([]string, error) {
	pm2Config, err := p.RetrievePm2Config()
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
	pm2Config, err := p.RetrievePm2Config()
	if err != nil {
		return nil, err
	}
	return pm2Config.ListHostsInEnv(env), nil
}

func (p *Pm2Manager) GetSSHConfig(env, host string) (*dto.SSHConfig, error) {
	pm2Config, err := p.RetrievePm2Config()
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
