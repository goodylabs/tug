package checkconnections

import (
	"fmt"
	"strings"
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

	hostToEnvs := make(map[string][]string)

	envs := projectConfig.GetAvailableEnvs()
	for _, envName := range envs {
		envCfg, _ := projectConfig.GetEnvConfig(envName)
		for _, host := range envCfg.Hosts {
			sshTarget := fmt.Sprintf("%s@%s", envCfg.User, host)
			hostToEnvs[sshTarget] = append(hostToEnvs[sshTarget], envName)
		}
	}

	var wg sync.WaitGroup
	statusTemplate := "  %-15s %-30s %s\n"

	for target, envList := range hostToEnvs {
		wg.Add(1)

		go func(sshTarget string, relatedEnvs []string) {
			defer wg.Done()

			parts := strings.SplitN(sshTarget, "@", 2)
			if len(parts) != 2 {
				return
			}
			user := parts[0]
			host := parts[1]

			sshCfg := &ports.SSHConfig{
				Host: host,
				Port: 22,
				User: user,
			}

			err := c.sshConnector.ConfigureSSHConnection(sshCfg)
			status := "✅ OK"
			if err != nil {
				status = "🚫 Connection failed"
			}
			for _, envName := range relatedEnvs {
				fmt.Printf(statusTemplate, envName, sshTarget, status)
			}
		}(target, envList)
	}

	wg.Wait()
	fmt.Println("\n✨ Verification finished.")
}
