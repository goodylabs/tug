package pm2

import (
	"encoding/json"
	"errors"
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

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s && read`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s && read`
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit %s`
	updateCmdTemplate   = `source ~/.nvm/nvm.sh; pm2 update`
)

var CommandTemplates = map[string]string{
	"[pm2] logs <resource>":     logsCmdTemplate,
	"[pm2] show <resource>":     showCmdTemplate,
	"[pm2] restart <resource>":  restartCmdTemplate,
	"[pm2] describe <resource>": describeCmdTemplate,
	"[pm2] monit <resource>":    monitCmdTemplate,
	"[pm2] update":              updateCmdTemplate,
}

const tmpJsonPath = "/tmp/ecosystem.json"

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

func (p *Pm2Manager) GetAvailableResources(sshConfig *dto.SSHConfig) ([]string, error) {
	output, err := p.sshConnector.RunCommand(jlistCmd)
	if err != nil {
		return nil, err
	}

	var pm2List []dto.Pm2ListItem
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		return nil, errors.New("failed to parse PM2 list output: " + err.Error())
	}

	var resources []string
	for _, item := range pm2List {
		resources = append(resources, item.Name)
	}

	return resources, nil
}

// NOT IN TECHNOLOGY HANDLER INTERFACE

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
