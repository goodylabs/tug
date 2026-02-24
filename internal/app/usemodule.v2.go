package app

import (
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/modules"
	"github.com/goodylabs/tug/internal/modules/action"
	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/goodylabs/tug/internal/ports"
)

type stepFunc func() (stepFunc, error)

type UseModuleV2UseCase struct {
	prompter      ports.Prompter
	sshService    *action.SSHService
	actionMgr     *action.ActionManager
	projectConfig modules.ProjectConfig

	ctx struct {
		env      string
		hostname string
		template string
	}
	stack []stepFunc
}

func NewUseModuleV2UseCase() *UseModuleV2UseCase {
	return &UseModuleV2UseCase{
		sshService: action.NewSSHService(adapters.NewSSHConnector()),
		prompter:   adapters.NewPrompter(),
	}
}

func (u *UseModuleV2UseCase) Execute(loadTech loadproject.StrategyName, actionTech action.StrategyName) error {
	lp := loadproject.NewLoadProject()
	pCfg, err := lp.Execute(loadTech)
	if err != nil {
		return err
	}
	u.projectConfig = pCfg

	actStrategy, err := action.GetStrategy(actionTech)
	if err != nil {
		return err
	}
	u.actionMgr = action.NewActionManager(actStrategy)

	u.stack = []stepFunc{u.stepSelectEnv}
	return u.runStack()
}

func (u *UseModuleV2UseCase) runStack() error {
	for len(u.stack) > 0 {
		stackLen := len(u.stack)
		currentStep := u.stack[stackLen-1]

		nextStep, err := currentStep()
		if err != nil {
			return err
		}

		if nextStep == nil {
			u.stack = u.stack[:stackLen-1]
			continue
		}

		if fmt.Sprintf("%v", nextStep) == fmt.Sprintf("%v", currentStep) {
			continue
		}

		u.stack = append(u.stack, nextStep)
	}
	return nil
}

func (u *UseModuleV2UseCase) stepSelectEnv() (stepFunc, error) {
	envs := u.projectConfig.GetAvailableEnvs()
	selected, err := u.prompter.ChooseFromList(envs, "Choose environment:")
	if err != nil {
		return nil, nil
	}

	u.ctx.env = selected
	return u.stepSelectHost, nil
}

func (u *UseModuleV2UseCase) stepSelectHost() (stepFunc, error) {
	hosts, err := u.projectConfig.GetAvailableHosts(u.ctx.env)
	if err != nil {
		return nil, err
	}

	label := fmt.Sprintf("[%s] Select host:", u.ctx.env)
	host, err := u.prompter.ChooseFromList(hosts, label)
	if err != nil {
		return nil, nil
	}

	envCfg, _ := u.projectConfig.GetEnvConfig(u.ctx.env)

	fmt.Println("Connecting...")
	hostname, err := u.sshService.Connect(envCfg.User, host)
	if err != nil {
		return nil, err
	}

	u.ctx.hostname = hostname
	return u.stepSelectAction, nil
}

func (u *UseModuleV2UseCase) stepSelectAction() (stepFunc, error) {
	templates := u.actionMgr.GetAvailableActionTemplates()

	label := fmt.Sprintf("[%s] Select action:", u.ctx.hostname)
	selected, err := u.prompter.ChooseFromMap(templates, label)
	if err != nil {
		return nil, nil
	}

	u.ctx.template = selected

	if strings.Contains(selected, "%s") {
		return u.stepSelectResource, nil
	}

	u.sshService.RunAction(u.ctx.template, "")
	return u.stepSelectAction, nil
}

func (u *UseModuleV2UseCase) stepSelectResource() (stepFunc, error) {
	res, err := u.actionMgr.GetAvailableResources(u.sshService.GetConnector())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch resources: %w", err)
	}

	label := fmt.Sprintf("[%s] Select resource:", u.ctx.hostname)
	resource, err := u.prompter.ChooseFromList(res, label)
	if err != nil {
		return nil, nil
	}

	u.sshService.RunAction(u.ctx.template, resource)
	return u.stepSelectResource, nil
}
