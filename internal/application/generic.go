package application

import (
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils/stageorchestrator"
)

type GenericUseCase struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
	context      struct {
		selectedEnv string
		sshConfig   *dto.SSHConfig
		action      string
		resource    string
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
	steps := []stageorchestrator.StepFunc{
		g.stepSelectEnv,
		g.stepSelectHost,
		g.stepSelectResource,
		g.stepSelectAction,
		g.stepExecuteAction,
	}

	stageOrchestrator := stageorchestrator.NewStageOrchestrator(steps)

	if err := g.handler.LoadConfigFromFile(); err != nil {
		return err
	}

	return stageOrchestrator.Run()
}

func (g *GenericUseCase) stepSelectEnv() (bool, error) {
	availableEnvs, err := g.handler.GetAvailableEnvs()
	if err != nil {
		return false, err
	}

	selectedEnv, err := g.prompter.ChooseFromList(availableEnvs, "ENVIRONMENTS")
	if err != nil {
		return false, nil
	}

	g.context.selectedEnv = selectedEnv

	return true, nil
}

func (g *GenericUseCase) stepSelectHost() (bool, error) {
	availableHosts, err := g.handler.GetAvailableHosts(g.context.selectedEnv)
	if err != nil {
		return false, err
	}

	selectedHost, err := g.prompter.ChooseFromList(availableHosts, "HOSTS")
	if err != nil {
		return false, nil
	}
	fmt.Println("Connecting to server...")

	sshConfig, err := g.handler.GetSSHConfig(g.context.selectedEnv, selectedHost)
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

	resource, err := g.prompter.ChooseFromList(resources, "RESOURCES")
	if err != nil {
		g.context.sshConfig = nil
		return false, nil
	}
	g.context.resource = resource
	return true, nil
}

func (g *GenericUseCase) stepSelectAction() (bool, error) {
	cmdTemplate, err := g.prompter.ChooseFromMap(g.handler.GetAvailableActionTemplates(), "ACTIONS")
	if err != nil {
		g.context.resource = ""
		fmt.Println("Looking for resources...")
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
