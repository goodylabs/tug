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
	monitCmdTemplate    = `source ~/.nvm/nvm.sh; pm2 monit %s`
)

var commandTemplates = map[string]string{
	"pm2 logs <resource>":     logsCmdTemplate,
	"pm2 show <resource>":     showCmdTemplate,
	"pm2 restart <resource>":  restartCmdTemplate,
	"pm2 describe <resource>": describeCmdTemplate,
	"pm2 monit <resource>":    monitCmdTemplate,
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

	var pm2List []dto.Pm2ListItem
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
	return p.prompter.ChooseFromList(resources, "Select PM2 <resource>")
}

func (p *Pm2Manager) SelectCommandTemplate() (string, error) {
	var commandNames []string
	for name := range commandTemplates {
		commandNames = append(commandNames, name)
	}

	selected, _ := p.prompter.ChooseFromList(commandNames, "Select command")

	cmdTemplate, ok := commandTemplates[selected]
	if !ok {
		return "", fmt.Errorf("selected command not found: %s", selected)
	}

	return cmdTemplate, nil
}
