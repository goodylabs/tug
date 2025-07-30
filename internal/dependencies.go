package internal

import (
	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/internal/services/docker"
	"github.com/goodylabs/tug/internal/services/pm2"
	"go.uber.org/dig"
)

func InitDependencyContainer() *dig.Container {
	container := dig.New()

	container.Provide(adapters.NewSSHConnector)
	container.Provide(adapters.NewPrompter)

	container.Provide(docker.NewDockerManager)
	container.Provide(pm2.NewPm2Manager)

	container.Provide(application.NewPm2UseCase)
	container.Provide(application.NewDockerUseCase)
	container.Provide(application.NewInitializeUseCase)

	return container
}
