package pm2

import (
	"fmt"
	"log"
	"os/exec"

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

func (p *Pm2Manager) LoadPm2Config(ecosystemConfigPath string, pm2Config *dto.EconsystemConfigDTO) error {
	tmpPath := `/tmp/ecosystem.json`

	script := fmt.Sprintf(`
		const fs = require("fs");
		const path = require("path");
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

	cmd := exec.Command("node", "-e", script)
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	err = utils.ReadJSON(tmpPath, &pm2Config)
	return err
}

func (p *Pm2Manager) SelectEnvironment(pm2Config *dto.EconsystemConfigDTO) string {
	if len(pm2Config.Deploy) == 0 {
		log.Fatal("No environments found in PM2 config")
	}

	var options []string
	for env := range pm2Config.Deploy {
		options = append(options, env)
	}

	return p.prompter.ChooseFromList(options, "Select pm2 environment")
}

func (p *Pm2Manager) selectHost(pm2Config *dto.EconsystemConfigDTO, selectedEnv string) string {
	if len(pm2Config.Deploy[selectedEnv].Host) == 0 {
		log.Fatal("No hosts found for the selected environment in PM2 config")
	}

	var options []string
	for _, host := range pm2Config.Deploy[selectedEnv].Host {
		options = append(options, host)
	}

	if len(options) == 1 {
		return options[0]
	}

	return p.prompter.ChooseFromList(options, "Select host for environment "+selectedEnv)
}

func (p *Pm2Manager) GetSSHConfig(pm2Config *dto.EconsystemConfigDTO, selectedEnv string) *dto.SSHConfigDTO {
	if _, exists := pm2Config.Deploy[selectedEnv]; !exists {
		log.Fatalf("Environment %s not found in PM2 config", selectedEnv)
	}

	sshConfig := dto.SSHConfigDTO{
		User: pm2Config.Deploy[selectedEnv].User,
		Host: p.selectHost(pm2Config, selectedEnv),
		Port: 22,
	}

	return &sshConfig
}
