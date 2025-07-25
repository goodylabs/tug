package mocks

import (
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/docker"
	"github.com/goodylabs/tug/internal/services/pm2"
	"go.uber.org/dig"
)

func InitTestingDependencyContainer(promptListSeq []int) *dig.Container {
	container := dig.New()

	container.Provide(NewSSHConnectorMock)

	container.Provide(func() ports.Prompter {
		return NewPrompterMock(promptListSeq)
	})

	container.Provide(docker.NewDockerManager)
	container.Provide(pm2.NewPm2Manager)

	container.Provide(application.NewPm2UseCase)
	container.Provide(application.NewDockerUseCase)

	return container
}
