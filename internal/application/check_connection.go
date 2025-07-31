package application

import (
	"fmt"

	"github.com/goodylabs/tug/internal/ports"
)

type CheckConnectionUseCase struct {
	handler      ports.TechnologyHandler
	sshConnector ports.SSHConnector
}

func NewCheckConnectionUseCase(handler ports.TechnologyHandler, sshConnector ports.SSHConnector) *CheckConnectionUseCase {
	return &CheckConnectionUseCase{
		handler:      handler,
		sshConnector: sshConnector,
	}
}

func (p *CheckConnectionUseCase) Execute() error {
	availableEnvs, err := p.handler.GetAvailableEnvs()
	if err != nil {
		return err
	}

	for _, env := range availableEnvs {
		hosts, err := p.handler.GetAvailableHosts(env)
		if err != nil {
			return err
		}

		for _, host := range hosts {
			template := "Env: %s, Host: %s - %s\n"
			sshConfig, err := p.handler.GetSSHConfig(env, host)
			if err != nil {
				return err
			}
			sshConStr := sshConfig.GetSSHConnectionString()
			if err := p.sshConnector.ConfigureSSHConnection(sshConfig); err != nil {
				fmt.Printf(template, env, sshConStr, "🚫 can not ssh connect")
				continue
			}

			_, err = p.handler.GetAvailableResources(sshConfig)
			if err != nil {
				fmt.Printf(template, env, sshConStr, "⚠️ can not list resources")
				continue
			}

			fmt.Printf(template, env, sshConStr, "✅ - everything is ok")
		}
	}
	// p.sshConnector.SetSSHConfig(selectedEnv)

	// resources, err := p.handler.GetAvailableResources(p.sshConnector.GetSSHConfig())
	// if err != nil {
	// 	return false, err
	// }

	// if len(resources) == 0 {
	// 	return false, nil // No resources available
	// }

	// selectedResource, err := prompter.SelectResource(resources)
	// if err != nil {
	// 	return false, err
	// }

	// p.sshConnector.SetResource(selectedResource)

	return nil
}
