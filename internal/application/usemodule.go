package application

import (
	"fmt"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils/stageorchestrator"
)

type UseModuleUseCase struct {
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

func NewUseModuleUseCase(handler ports.TechnologyHandler, sshConnector ports.SSHConnector, prompter ports.Prompter) *UseModuleUseCase {
	return &UseModuleUseCase{
		handler:      handler,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (u *UseModuleUseCase) Execute() error {
	steps := []stageorchestrator.StepFunc{
		u.stepSelectEnv,
		u.stepSelectHost,
		u.stepSelectResource,
		u.stepSelectAction,
		u.stepExecuteAction,
	}

	stageOrchestrator := stageorchestrator.NewStageOrchestrator(steps)

	if err := u.handler.LoadConfigFromFile(); err != nil {
		return err
	}

	return stageOrchestrator.Run()
}

func (u *UseModuleUseCase) stepSelectEnv() (bool, error) {
	availableEnvs, err := u.handler.GetAvailableEnvs()
	if err != nil {
		return false, err
	}

	selectedEnv, err := u.prompter.ChooseFromList(availableEnvs, "ENVIRONMENTS")
	if err != nil {
		return false, nil
	}

	u.context.selectedEnv = selectedEnv

	return true, nil
}

func (u *UseModuleUseCase) stepSelectHost() (bool, error) {
	availableHosts, err := u.handler.GetAvailableHosts(u.context.selectedEnv)
	if err != nil {
		return false, err
	}

	selectedHost, err := u.prompter.ChooseFromList(availableHosts, "HOSTS")
	if err != nil {
		return false, nil
	}
	fmt.Println("Connecting to server...")

	sshConfig, err := u.handler.GetSSHConfig(u.context.selectedEnv, selectedHost)
	if err != nil {
		return false, err
	}

	u.sshConnector.ConfigureSSHConnection(sshConfig)
	u.context.sshConfig = sshConfig

	return true, nil
}

func (u *UseModuleUseCase) stepSelectResource() (bool, error) {
	resources, err := u.handler.GetAvailableResources(u.context.sshConfig)
	if err != nil {
		return false, err
	}

	resource, err := u.prompter.ChooseFromList(resources, "RESOURCES")
	if err != nil {
		u.context.sshConfig = nil
		return false, nil
	}
	u.context.resource = resource
	return true, nil
}

func (u *UseModuleUseCase) stepSelectAction() (bool, error) {
	cmdTemplate, err := u.prompter.ChooseFromMap(u.handler.GetAvailableActionTemplates(), "ACTIONS")
	if err != nil {
		u.context.resource = ""
		fmt.Println("Looking for resources...")
		return false, nil
	}
	u.context.action = cmdTemplate
	return true, nil
}

func (u *UseModuleUseCase) stepExecuteAction() (bool, error) {
	remoteCmd := fmt.Sprintf(u.context.action, u.context.resource)
	if err := u.sshConnector.RunInteractiveCommand(remoteCmd); err != nil {
		u.context.action = ""
	}
	return false, nil
}
