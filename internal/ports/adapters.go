package ports

import "github.com/goodylabs/docker-swarm-cli/internal/dto"

type PromptPort interface {
	ChooseFromList(options []string, label string) (string, error)
}

type DockerPort interface {
	ListContainers() []dto.ContainerDTO
	ConfigureDocker(string)
}
