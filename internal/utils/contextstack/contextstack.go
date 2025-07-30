package contextstack

import (
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/pm2"
)

type Generic struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
	sshConfig    *dto.SSHConfig
	action       string
	resource     string
}

func NewGeneric(handler ports.TechnologyHandler, sshConnector ports.SSHConnector, prompter ports.Prompter) *Generic {
	return &Generic{
		handler:      handler,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (g *Generic) Execute() error {
	for true {
		if g.sshConfig == nil {
			availableEnvs, err := g.handler.GetAvailableEnvs()
			if err != nil {
				return fmt.Errorf("selecting PM2 environment: %w", err)
			}

			selectedEnv, err := g.prompter.ChooseFromList(availableEnvs, "Select PM2 <environment>")
			if err != nil {
				return fmt.Errorf("selecting PM2 environment: %w", err)
			}

			sshConfig, err := g.handler.GetSSHConfig(selectedEnv)
			if err != nil {
				return fmt.Errorf("getting SSH config: %w", err)
			}

			g.sshConnector.ConfigureSSHConnection(sshConfig)

			g.sshConfig = sshConfig
			continue
		}

		if g.resource == "" {
			resources, err := g.handler.GetAvailableResources(g.sshConfig)
			if err != nil {
				return fmt.Errorf("selecting PM2 resource: %w", err)
			}

			resource, err := g.prompter.ChooseFromList(resources, "Select PM2 <resource>")
			if err != nil {
				g.sshConfig = nil
				continue
			}
			g.resource = resource
			continue
		}

		if g.action == "" {
			cmdTemplate, err := g.prompter.ChooseFromMap(pm2.CommandTemplates, "Chose command")
			if err != nil {
				g.resource = ""
				continue
			}
			g.action = cmdTemplate
			continue
		}

		remoteCmd := fmt.Sprintf(g.action, g.resource)
		if err := g.sshConnector.RunInteractiveCommand(remoteCmd); err != nil {
			g.action = ""
			continue
		}
	}
	return nil
}
