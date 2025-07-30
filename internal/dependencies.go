package internal

import (
	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/internal/services/docker"
	"github.com/goodylabs/tug/internal/services/pm2"
	"go.uber.org/dig"
)

type managerType string

const (
	DockerManager managerType = "docker"
	Pm2Manager    managerType = "pm2"
)

func InitDependencyContainer(manager managerType) *dig.Container {
	container := dig.New()

	container.Provide(adapters.NewSSHConnector)
	container.Provide(adapters.NewPrompter)

	switch manager {
	case DockerManager:
		container.Provide(docker.NewDockerManager)
	case Pm2Manager:
		container.Provide(pm2.NewPm2Manager)
	}

	container.Provide(application.NewGenericUseCase)

	return container
}
