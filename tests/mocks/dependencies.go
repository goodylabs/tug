package mocks

import (
	"github.com/goodylabs/tug/internal/services/docker"
	"github.com/goodylabs/tug/internal/services/pm2"
)

func SetupMockPm2Manager(prompts []int, sshOutput string, sshErr error) *pm2.Pm2Manager {
	return pm2.NewPm2Manager(
		NewPrompterMock(prompts),
		NewSSHConnectorMock(sshOutput, sshErr),
	)
}

func SetupMockDockerManager(prompts []int, sshOutput string, sshErr error) *docker.DockerManager {
	return docker.NewDockerManager(
		NewPrompterMock(prompts),
		NewSSHConnectorMock(sshOutput, sshErr),
	)
}
