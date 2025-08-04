package internal

import (
	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/internal/modules/docker"
	"github.com/goodylabs/tug/internal/modules/pm2"
	"go.uber.org/dig"
)

type OptFunc func(*dig.Container)

func WithDockerHandler(container *dig.Container) {
	container.Provide(docker.NewDockerManager)
}

func WithPm2Handler(container *dig.Container) {
	container.Provide(pm2.NewPm2Handler)
}

func InitDependencyContainer(opts ...OptFunc) *dig.Container {
	container := dig.New()

	container.Provide(adapters.NewSSHConnector)
	container.Provide(adapters.NewPrompter)

	for _, opt := range opts {
		opt(container)
	}

	container.Provide(application.NewUseModuleUseCase)
	container.Provide(application.NewCheckConnectionUseCase)
	container.Provide(application.NewInitializeUseCase)

	return container
}
