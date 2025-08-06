package application

import (
	"fmt"
	"sync"

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
	if err := p.handler.LoadConfigFromFile(); err != nil {
		return err
	}

	availableEnvs, err := p.handler.GetAvailableEnvs()
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, 1)

	for _, env := range availableEnvs {
		hosts, err := p.handler.GetAvailableHosts(env)
		if err != nil {
			return err
		}

		for _, host := range hosts {
			wg.Add(1)
			go func(env, host string) {
				defer wg.Done()
				if err := p.checkHost(env, host); err != nil {
					select {
					case errCh <- err:
					default:
					}
				}
			}(env, host)
		}
	}

	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

func (p *CheckConnectionUseCase) checkHost(env, host string) error {
	template := "Env: %s, Host: %s - %s\n"

	sshConfig, err := p.handler.GetSSHConfig(env, host)
	if err != nil {
		return err
	}

	sshConStr := sshConfig.User + "@" + sshConfig.Host

	if err := p.sshConnector.ConfigureSSHConnection(sshConfig); err != nil {
		fmt.Printf(template, env, sshConStr, "🚫 can not ssh connect")
		return nil
	}

	if _, err := p.handler.GetAvailableResources(sshConfig); err != nil {
		fmt.Printf(template, env, sshConStr, "⚠️ can not list resources")
		return nil
	}

	fmt.Printf(template, env, sshConStr, "✅ - everything is ok")
	return nil
}
