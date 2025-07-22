package ports

import "github.com/goodylabs/docker-swarm-cli/internal/dto"

type PromptPort interface {
	ChooseFromList(options []string, label string) string
}

type DockerPort interface {
	ConfigureDocker(string)
	ListServices() []dto.ServiceDTO
	ListContainers() []dto.ContainerDTO
}
