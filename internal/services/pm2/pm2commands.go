package pm2

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
)

const (
	resourcesOption = "<resource>"
)

type commandOption struct {
	name    string
	command string
}

const (
	// jlistCmd = `pm2 jlist | sed -n '/^\[/,$p'`
	jlistCmd = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`
	logsCmd  = `source ~/.nvm/nvm.sh; pm2 logs <resource>`
)

func (p *Pm2Manager) SelectResource() string {
	output, err := p.sshConnector.RunCommand(jlistCmd)
	if err != nil {
		log.Fatalf("Error running PM2 command '%s' error: %v", jlistCmd, err)
	}

	var pm2List []dto.Pm2ListItemDTO
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		log.Fatalf("Error parsing PM2 list output: %v after running `%s`", err, jlistCmd)
	}

	var pm2Resources []string

	for _, item := range pm2List {
		pm2Resources = append(pm2Resources, item.Name)
	}

	pm2Resources = append(pm2Resources, constants.ALL_OPTION)

	return p.prompter.ChooseFromList(pm2Resources, "Select PM2 resource")
}

func (p *Pm2Manager) getCommandOptions(resource string) *[]commandOption {
	var selectedResource string

	switch resource {
	case constants.ALL_OPTION:
		selectedResource = ""
	default:
		selectedResource = resource
	}

	var commandOptions = []commandOption{
		{
			name:    "PM2 logs",
			command: strings.ReplaceAll(logsCmd, resourcesOption, selectedResource),
		},
	}

	return &commandOptions
}

func (p *Pm2Manager) RunCommandOnResource(resource string) {
	options := p.getCommandOptions(resource)

	commandNames := make([]string, len(*options))
	for i, option := range *options {
		commandNames[i] = option.name
	}
	selectedCommandName := p.prompter.ChooseFromList(commandNames, "Select command:")
	var selectedCommand string
	for _, option := range *options {
		if option.name == selectedCommandName {
			selectedCommand = option.command
			break
		}
	}

	if err := p.sshConnector.RunInteractiveCommand(selectedCommand); err != nil {
		log.Fatalf("Error running command on PM2 resource: %v", err)
	}
}
