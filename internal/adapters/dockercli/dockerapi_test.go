package dockercli_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters/dockercli"
	"github.com/goodylabs/tug/tests/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
)

var (
	containerName  = "redis"
	redisContainer testcontainers.Container
)

func init() {
	redisContainer = testutils.StartContainer(nil, testcontainers.ContainerRequest{
		Image: "redis:latest",
		Name:  containerName,
	})
}

func TestListingContainers(t *testing.T) {
	dockerApi := dockercli.NewDockerApi()
	containers := dockerApi.ListContainers()

	var found = false
	for _, container := range containers {
		if container.Name == containerName {
			found = true
			break
		}
	}

	assert.True(t, found)
}
