package testadapters

import (
	"testing"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/constants"
	testutils "github.com/goodylabs/docker-swarm-cli/tests/utils"
	"github.com/testcontainers/testcontainers-go"
)

func TestWithRedis(t *testing.T) {
	testutils.StartContainer(t, testcontainers.ContainerRequest{
		Image: "redis:latest",
	})

	dockerAdapter := adapters.NewDockerAdapter()
	dockerAdapter.ConfigureDocker(constants.LOCAL_DOCKER_HOST)

	dockerAdapter.ListContainers()
}
