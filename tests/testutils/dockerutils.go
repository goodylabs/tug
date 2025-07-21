package testutils

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
)

func StartContainer(t *testing.T, req testcontainers.ContainerRequest) testcontainers.Container {
	ctx := context.Background()

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if t != nil {
		require.NoError(t, err)
		t.Cleanup(func() {
			_ = container.Terminate(ctx)
		})
	} else if err != nil {
		panic(err)
	}

	return container
}
