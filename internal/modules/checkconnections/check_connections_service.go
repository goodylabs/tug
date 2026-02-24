package checkconnections

import (
	"fmt"
	"sync"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/modules"
	"github.com/goodylabs/tug/internal/ports"
)

type CheckConnectionsService struct {
	sshConnector ports.SSHConnector
}

func NewCheckConnectionsService() *CheckConnectionsService {
	return &CheckConnectionsService{
		sshConnector: adapters.NewSSHConnector(),
	}
}

func (c *CheckConnectionsService) Execute(projectConfig modules.ProjectConfig) {
	fmt.Printf("🔍 Verification for project config\n\n")

	var wg sync.WaitGroup
	envs := projectConfig.GetAvailableEnvs()

	for _, envName := range envs {
		envCfg, _ := projectConfig.GetEnvConfig(envName)

		for _, host := range envCfg.Hosts {
			wg.Add(1)

			go func(envName, host, user string) {
				defer wg.Done()

				statusTemplate := "  %-15s %-30s %s\n"
				sshTarget := fmt.Sprintf("%s@%s", user, host)

				sshCfg := &ports.SSHConfig{
					Host: host,
					Port: 22,
					User: user,
				}

				if err := c.sshConnector.ConfigureSSHConnection(sshCfg); err != nil {
					fmt.Printf(statusTemplate, envName, sshTarget, "🚫 Connection failed")
					return
				}

				fmt.Printf(statusTemplate, envName, sshTarget, "✅ OK")

			}(envName, host, envCfg.User)
		}
	}

	wg.Wait()
	fmt.Println("\n✨ Verification finished.")
}
