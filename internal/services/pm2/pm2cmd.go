package pm2

import (
	"encoding/json"

	"github.com/goodylabs/tug/internal/dto"
)

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s && read`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s && read`
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit %s && read`
)

var CommandTemplates = map[string]string{
	"pm2 logs <resource>":     logsCmdTemplate,
	"pm2 show <resource>":     showCmdTemplate,
	"pm2 restart <resource>":  restartCmdTemplate,
	"pm2 describe <resource>": describeCmdTemplate,
	"pm2 monit <resource>":    monitCmdTemplate,
}

func (p *Pm2Manager) GetAvailableResources(sshConfig *dto.SSHConfig) ([]string, error) {
	output, err := p.sshConnector.RunCommand(jlistCmd)
	if err != nil {
		return nil, err
	}

	var pm2List []dto.Pm2ListItem
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		return nil, err
	}

	var resources []string
	for _, item := range pm2List {
		resources = append(resources, item.Name)
	}

	return resources, nil
}
