package pm2

import (
	"encoding/json"
	"fmt"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
)

const (
	resourcesOption = "<resource>"
	jlistCmd        = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmdTemplate = `source ~/.nvm/nvm.sh; pm2 logs %s`
)

type commandOption struct {
	Name    string
	Command string
}

func (p *Pm2Manager) GetPm2Processes() ([]string, error) {
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

	resources = append(resources, constants.ALL_OPTION)

	return resources, nil
}

func (p *Pm2Manager) SelectResource() (string, error) {
	resources, err := p.GetPm2Processes()
	if err != nil {
		return "", err
	}
	return p.prompter.ChooseFromList(resources, "Select PM2 resource"), nil
}

func (p *Pm2Manager) GetCommandOptions(resource string) []commandOption {
	selectedResource := ""
	if resource != constants.ALL_OPTION {
		selectedResource = resource
	}

	return []commandOption{
		{
			Name:    "PM2 logs",
			Command: fmt.Sprintf(logsCmdTemplate, selectedResource),
		},
	}
}

func (p *Pm2Manager) RunCommandOnResource(resource string) error {
	options := p.GetCommandOptions(resource)

	var commandNames []string
	for _, opt := range options {
		commandNames = append(commandNames, opt.Name)
	}

	selectedCmdName := p.prompter.ChooseFromList(commandNames, "Select command:")
	for _, opt := range options {
		if opt.Name == selectedCmdName {
			return p.sshConnector.RunInteractiveCommand(opt.Command)
		}
	}

	return fmt.Errorf("selected command not found: %s", selectedCmdName)
}
