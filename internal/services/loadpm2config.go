package services

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/utils"
)

func GetLoadPm2Config(ecosystemConfigPath string) (*dto.EconsystemConfigDTO, error) {
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

	var pm2ConfigDTO dto.EconsystemConfigDTO

	err = utils.ReadJSON(tmpPath, &pm2ConfigDTO)
	return &pm2ConfigDTO, err
}
