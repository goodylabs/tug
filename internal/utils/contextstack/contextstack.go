package contextstack

import (
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/pm2"
)

type ContextStack struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
	sshConfig    *dto.SSHConfig
	action       string
	resource     string
}

func NewContextStack(handler ports.TechnologyHandler, sshConnector ports.SSHConnector, prompter ports.Prompter) *ContextStack {
	return &ContextStack{
		handler:      handler,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (c *ContextStack) Execute() error {
	for true {
		if c.sshConfig == nil {
			availableEnvs, err := c.handler.GetAvailableEnvs()
			if err != nil {
				return fmt.Errorf("selecting PM2 environment: %w", err)
			}

			selectedEnv, err := c.prompter.ChooseFromList(availableEnvs, "Select PM2 <environment>")
			if err != nil {
				return fmt.Errorf("selecting PM2 environment: %w", err)
			}

			sshConfig, err := c.handler.GetSSHConfig(selectedEnv)
			if err != nil {
				return fmt.Errorf("getting SSH config: %w", err)
			}

			c.sshConnector.ConfigureSSHConnection(sshConfig)

			c.sshConfig = sshConfig
			continue
		}

		if c.resource == "" {
			resources, err := c.handler.GetAvailableResources(c.sshConfig)
			if err != nil {
				return fmt.Errorf("selecting PM2 resource: %w", err)
			}

			resource, err := c.prompter.ChooseFromList(resources, "Select PM2 <resource>")
			if err != nil {
				c.sshConfig = nil
				continue
			}
			c.resource = resource
			continue
		}

		if c.action == "" {
			cmdTemplate, err := c.prompter.ChooseFromMap(pm2.CommandTemplates, "Chose command")
			if err != nil {
				c.resource = ""
				continue
			}
			c.action = cmdTemplate
			continue
		}

		remoteCmd := fmt.Sprintf(c.action, c.resource)
		if err := c.sshConnector.RunInteractiveCommand(remoteCmd); err != nil {
			c.action = ""
			continue
		}
	}
	return nil
}
