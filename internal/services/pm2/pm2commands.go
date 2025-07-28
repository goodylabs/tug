package pm2

import (
	"encoding/json"
	"log"

	"github.com/goodylabs/tug/internal/dto"
)

const (
	JLIST_CMD = `pm2 jlist | sed -n '/^\[/,$p'`
)

func (p *Pm2Manager) jsonOutputHandler(output string, dtoStruct any) error {
	return json.Unmarshal([]byte(output), &dtoStruct)
}

func (p *Pm2Manager) SelectResource() string {
	output, err := p.sshConnector.RunCommand(JLIST_CMD)
	if err != nil {
		log.Fatalf("Error running PM2 command: %v", err)
	}

	var pm2List []dto.Pm2ListItemDTO
	err = p.jsonOutputHandler(output, &pm2List)
	if err != nil {
		log.Fatalf("Error parsing PM2 list output: %v", err)
	}
	return pm2List[0].Name
}
