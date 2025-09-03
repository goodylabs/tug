package services_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/modules/docker/services"
	"github.com/stretchr/testify/assert"
)

func TestGetResourcesFromJsonOutput(t *testing.T) {
	output := `
	{"Command":"\"/entrypoint.sh --co…\"","CreatedAt":"2025-09-03 08:01:25 +0000 UTC","ID":"3f03d2082d78","Image":"traefik:v3.0","Names":"traefik_traefik.ue8","Size":"0B","State":"running","Status":"Up 2 hours"}
	{"Command":"\"./docker-entrypoint…\"","CreatedAt":"2025-09-03 04:57:47 +0000 UTC","ID":"59a393c15ece","Image":"ghcr.io/some/test:654321","Names":"test_test.123","Size":"0B","State":"running","Status":"Up 5 hours (healthy)"}
	{"Command":"\"/usr/bin/vector\"","CreatedAt":"2025-08-28 09:32:38 +0000 UTC","ID":"b1aeaf771207","Image":"timberio/vector:0.47.0-debian","Names":"vector-vector-1","Size":"0B","State":"running","Status":"Up 6 days"}
	`
	containerNames, err := services.GetResourcesFromJsonOutput(output)
	assert.NoError(t, err)
	assert.Equal(t, []string{"traefik_traefik.ue8", "test_test.123", "vector-vector-1"}, containerNames)
}
