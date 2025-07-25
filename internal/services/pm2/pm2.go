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
}

func NewPm2Manager() ports.Pm2Manager {
	return &Pm2Manager{}
}

func (p *Pm2Manager) LoadPm2Config(ecosystemConfigPath string, pm2ConfigDTO *dto.EconsystemConfigDTO) error {
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

	err = utils.ReadJSON(tmpPath, &pm2ConfigDTO)
	return err
}
