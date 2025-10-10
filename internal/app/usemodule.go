package app

import (
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/utils"
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

	selectedEnv, err := u.prompter.ChooseFromList(availableEnvs, "Choose an environment:")
	if err != nil {
		return nil, nil
	}

	u.context.selectedEnv = selectedEnv

	return u.stepSelectHost, nil
}

const helpCmdTemplate = `
1. Check if you can connect as root manually: ssh %s
2. If you can, please run the following command to copy your user's SSH keys from root user to app user on host:

grep -vxFf %s /root/.ssh/authorized_keys >> %s

3. Try again to connect with tug.
`

func (u *UseModuleUseCase) stepSelectHost() (stepFunc, error) {
	availableHosts, err := u.handler.GetAvailableHosts(u.context.selectedEnv)
	if err != nil {
		return nil, err
	}

	promptLabel := fmt.Sprintf("[ %s ] Choose a host:", u.context.selectedEnv)
	selectedHost, err := u.prompter.ChooseFromList(availableHosts, promptLabel)
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
		if sshConfig.User == "root" {
			return nil, fmt.Errorf("Failed to connect to the server with %s - error: %s", userAddr, err.Error())
		}
		rootAddr := "root" + "@" + sshConfig.Host
		userAuthKeys := fmt.Sprintf("/home/%s/.ssh/authorized_keys", sshConfig.User)
		helpCommand := fmt.Sprintf(helpCmdTemplate, rootAddr, userAuthKeys, userAuthKeys)
		return nil, fmt.Errorf("Failed to connect to the server with %s - error: %s\nWhy is that? \n%s\n", userAddr, err.Error(), helpCommand)
	}

	u.context.sshConfig = sshConfig

	return u.stepSelectAction, nil
}

func (u *UseModuleUseCase) getRemoteHostname() string {
	output, err := u.sshConnector.RunCommand("hostname")
	if err != nil {
		return "failed_to_display_hostname"
	}
	return output
}

func (u *UseModuleUseCase) stepSelectAction() (stepFunc, error) {
	actionTemplates := u.handler.GetAvailableActionTemplates()

	remoteHostname := u.getRemoteHostname()
	promptLabel := fmt.Sprintf("[ %s ] Choose an action:", utils.NormalizeSpaces(remoteHostname))
	fmt.Println(promptLabel)
	cmdTemplate, err := u.prompter.ChooseFromMap(actionTemplates, promptLabel)
	if err != nil {
		fmt.Println("Moving back to host selection...")
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

	actionName := utils.NormalizeSpaces(u.context.action)
	promptLabel := fmt.Sprintf("[ %s ] Choose a resource:", actionName)
	resource, err := u.prompter.ChooseFromList(resources, promptLabel)
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
