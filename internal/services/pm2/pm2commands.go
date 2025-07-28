package pm2

import (
	"encoding/json"
	"log"

	"github.com/goodylabs/tug/internal/dto"
)

const (
	JLIST_CMD = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
)

func (p *Pm2Manager) SelectResource() string {
	output, err := p.sshConnector.RunCommand(JLIST_CMD)

	log.Println("Raw PM2 output:")
	log.Println(output)

	if err != nil {
		log.Fatalf("Error running PM2 command: %v", err)
	}

	var pm2List []dto.Pm2ListItemDTO
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		log.Fatalf("Error parsing PM2 list output: %v after running `%s`", err, JLIST_CMD)
	}

	pm2Resources := make([]string, len(pm2List))
	for i, item := range pm2List {
		pm2Resources[i] = item.Name
	}
	if len(pm2Resources) == 0 {
		log.Fatal("No PM2 resources found")
	}

	return p.prompter.ChooseFromList(pm2Resources, "Select PM2 resource")
}
