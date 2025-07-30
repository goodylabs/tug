package pm2

import (
	"encoding/json"
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
)

const (
	jlistCmd            = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 logs %s`
	showCmdTemplate     = `source ~/.nvm/nvm.sh; pm2 show %s`
	restartCmdTemplate  = `source ~/.nvm/nvm.sh; pm2 restart %s`
	describeCmdTemplate = `source ~/.nvm/nvm.sh; pm2 describe %s`
)

var commandTemplates = map[string]string{
	"pm2 logs":        logsCmdTemplate,
	"pm2 show status": showCmdTemplate,
	"pm2 restart":     restartCmdTemplate,
	"pm2 describe":    describeCmdTemplate,
}

type commandOption struct {
	Name            string
	CommandTemplate string
}

func (p *Pm2Manager) getPm2Processes() ([]string, error) {
	output, err := p.sshConnector.RunCommand(jlistCmd)
	if err != nil {
		return nil, fmt.Errorf("running PM2 jlist command: %w", err)
	}

	var pm2List []dto.Pm2ListItemDTO
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		return nil, fmt.Errorf("parsing PM2 list output: %w", err)
	}

	var resources []string
	for _, item := range pm2List {
		resources = append(resources, item.Name)
	}

	return resources, nil
}

func (p *Pm2Manager) SelectResource() (string, error) {
	resources, err := p.getPm2Processes()
	if err != nil {
		return "", err
	}
	return p.prompter.ChooseFromList(resources, "Select PM2 resource"), nil
}

func (p *Pm2Manager) SelectCommandTemplate() (string, error) {
	var commandNames []string
	for name := range commandTemplates {
		commandNames = append(commandNames, name)
	}

	selected := p.prompter.ChooseFromList(commandNames, "Select command:")

	cmdTemplate, ok := commandTemplates[selected]
	if !ok {
		return "", fmt.Errorf("selected command not found: %s", selected)
	}

	return cmdTemplate, nil
}
