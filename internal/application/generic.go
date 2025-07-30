package application

import (
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/pm2"
)

type GenericUseCase struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
	context      struct {
		sshConfig *dto.SSHConfig
		action    string
		resource  string
	}
}

func NewGenericUseCase(handler ports.TechnologyHandler, sshConnector ports.SSHConnector, prompter ports.Prompter) *GenericUseCase {
	return &GenericUseCase{
		handler:      handler,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (g *GenericUseCase) Execute() error {
	steps := []func() (bool, error){
		g.stepSelectEnv,
		g.stepSelectResource,
		g.stepSelectAction,
		g.stepExecuteAction,
	}

	currentStep := 0

	for currentStep < len(steps) {
		nextStep, err := steps[currentStep]()
		if err != nil {
			return err
		}
		if nextStep {
			currentStep++
		} else {
			if currentStep == 0 {
				fmt.Println("Exiting PM2 command execution.")
				return nil
			}
			currentStep--
		}

	}
	return nil
}

func (g *GenericUseCase) stepSelectEnv() (bool, error) {
	availableEnvs, err := g.handler.GetAvailableEnvs()
	if err != nil {
		return false, err
	}

	selectedEnv, err := g.prompter.ChooseFromList(availableEnvs, "Select PM2 <environment>")
	if err != nil {
		return false, nil
	}

	sshConfig, err := g.handler.GetSSHConfig(selectedEnv)
	if err != nil {
		return false, err
	}

	g.sshConnector.ConfigureSSHConnection(sshConfig)
	g.context.sshConfig = sshConfig

	return true, nil
}

func (g *GenericUseCase) stepSelectResource() (bool, error) {
	resources, err := g.handler.GetAvailableResources(g.context.sshConfig)
	if err != nil {
		return false, err
	}

	resource, err := g.prompter.ChooseFromList(resources, "Select PM2 <resource>")
	if err != nil {
		g.context.sshConfig = nil
		return false, nil
	}
	g.context.resource = resource
	return true, nil
}

func (g *GenericUseCase) stepSelectAction() (bool, error) {
	cmdTemplate, err := g.prompter.ChooseFromMap(pm2.CommandTemplates, "Chose command")
	if err != nil {
		g.context.resource = ""
		return false, nil
	}
	g.context.action = cmdTemplate
	return true, nil
}

func (g *GenericUseCase) stepExecuteAction() (bool, error) {
	remoteCmd := fmt.Sprintf(g.context.action, g.context.resource)
	if err := g.sshConnector.RunInteractiveCommand(remoteCmd); err != nil {
		g.context.action = ""
	}
	return false, nil
}
