package app

import (
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
)

type stepFunc func() (stepFunc, error)

type UseModuleUseCase struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
	context      struct {
		selectedEnv string
		sshConfig   *ports.SSHConfig
		action      string
		resource    string
	}
	stack []stepFunc
}

func NewUseModuleUseCase(handler ports.TechnologyHandler, sshConnector ports.SSHConnector, prompter ports.Prompter) *UseModuleUseCase {
	return &UseModuleUseCase{
		handler:      handler,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (u *UseModuleUseCase) Execute() error {

	if err := u.handler.LoadConfigFromFile(); err != nil {
		return err
	}

	u.stack = []stepFunc{u.stepSelectEnv}

	for len(u.stack) > 0 {
		stackLen := len(u.stack)
		nextStep, err := u.stack[stackLen-1]()
		if err != nil {
			return err
		}
		if nextStep == nil {
			u.stack = u.stack[:stackLen-1]
			continue
		}
		u.stack = append(u.stack, nextStep)
	}

	return nil
}

func (u *UseModuleUseCase) stepSelectEnv() (stepFunc, error) {
	availableEnvs, err := u.handler.GetAvailableEnvs()
	if err != nil {
		return nil, err
	}

	selectedEnv, err := u.prompter.ChooseFromList(availableEnvs, "ENVIRONMENTS")
	if err != nil {
		return nil, nil
	}

	u.context.selectedEnv = selectedEnv

	return u.stepSelectHost, nil
}

func (u *UseModuleUseCase) stepSelectHost() (stepFunc, error) {
	availableHosts, err := u.handler.GetAvailableHosts(u.context.selectedEnv)
	if err != nil {
		return nil, err
	}

	selectedHost, err := u.prompter.ChooseFromList(availableHosts, "HOSTS")
	if err != nil {
		return nil, nil
	}
	fmt.Println("Connecting to server...")

	sshConfig, err := u.handler.GetSSHConfig(u.context.selectedEnv, selectedHost)
	if err != nil {
		return nil, err
	}

	if err = u.sshConnector.ConfigureSSHConnection(sshConfig); err != nil {
		userAddr := sshConfig.User + "@" + sshConfig.Host
		if sshConfig.User != "root" {
			return nil, fmt.Errorf("Failed to connect to the server with %s", userAddr)
		}
		rootAddr := "root" + "@" + sshConfig.Host
		userAuthKeys := fmt.Sprintf("/home/%s/.ssh/authorized_keys", sshConfig.User)
		helpCommand := fmt.Sprintf("ssh %s grep -vxFf %s /root/.ssh/authorized_keys >> %s", rootAddr, userAuthKeys, userAuthKeys)
		return nil, fmt.Errorf("Failed to connect to the server with %s - probably the user doesn't has your ssh key.\nYou can fix it with command: \n\n%s\n", userAddr, helpCommand)
	}
	u.context.sshConfig = sshConfig

	return u.stepSelectAction, nil
}

func (u *UseModuleUseCase) stepSelectAction() (stepFunc, error) {
	cmdTemplate, err := u.prompter.ChooseFromMap(u.handler.GetAvailableActionTemplates(), "ACTIONS")
	if err != nil {
		fmt.Println("Looking for resources...")
		return nil, nil
	}

	u.context.action = cmdTemplate

	if strings.Contains(cmdTemplate, "%s") {
		return u.stepSelectResource, nil
	} else {
		return u.stepExecuteAction, nil
	}
}

func (u *UseModuleUseCase) stepSelectResource() (stepFunc, error) {
	resources, err := u.handler.GetAvailableResources(u.context.sshConfig)
	if err != nil {
		return nil, err
	}

	resource, err := u.prompter.ChooseFromList(resources, "RESOURCES")
	if err != nil {
		u.context.sshConfig = nil
		return nil, nil
	}
	u.context.resource = resource
	return u.stepExecuteAction, nil
}

func (u *UseModuleUseCase) stepExecuteAction() (stepFunc, error) {
	var remoteCmd = u.context.action

	if strings.Contains(remoteCmd, "%s") {
		remoteCmd = fmt.Sprintf(u.context.action, u.context.resource)
	}

	u.sshConnector.RunInteractiveCommand(remoteCmd)
	return nil, nil
}
