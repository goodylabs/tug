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

func NewPm2Manager(prompter ports.Prompter, sshconnector ports.SSHConnector) *Pm2Manager {
	return &Pm2Manager{
		prompter:     prompter,
		sshConnector: sshconnector,
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

	// p.prompter.Print("Available environments:")
	return "staging"
}
