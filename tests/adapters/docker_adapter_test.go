package testadapters

import (
	"testing"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/constants"
	testutils "github.com/goodylabs/docker-swarm-cli/tests/utils"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

func TestWithRedis(t *testing.T) {
	testutils.StartContainer(t, testcontainers.ContainerRequest{
		Image: "redis:latest",
		Name:  "redis",
	})

	dockerAdapter := adapters.NewDockerAdapter()
	dockerAdapter.ConfigureDocker(constants.LOCAL_DOCKER_HOST)

	containers := dockerAdapter.ListContainers()
	isRedis := false
	for _, container := range containers {
		if container.Name == "redis" {
			isRedis = true
		}
	}
	assert.True(t, isRedis)
}
