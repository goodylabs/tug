package adapters

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
)

var dockerClient *client.Client

type DockerAdapter struct{}

func (d *DockerAdapter) ConfigureDocker(dockerHost string) {
	var err error
	dockerClient, err = client.NewClientWithOpts(
		client.WithHost(dockerHost),
		client.WithVersion(config.DOCKER_API_VERSION),
	)
	if err != nil {
		panic(err)
	}
}

func (d *DockerAdapter) ListContainers() []dto.ContainerDTO {
	containers, err := dockerClient.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	containerDTOs := []dto.ContainerDTO{}
	for _, container := range containers {
		containerDTOs = append(containerDTOs, dto.ContainerDTO{
			Id:   container.ID,
			Name: container.Names[0][1:],
		})
	}

	return containerDTOs
}

func NewDockerAdapter() ports.DockerPort {
	return &DockerAdapter{}
}
