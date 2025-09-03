package services_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/modules/docker/swarm/services"
	"github.com/stretchr/testify/assert"
)

func TestGetResourcesFromJsonOutput(t *testing.T) {
	output := `
	{"ID":"vkxz3t57kouf","Image":"traefik:v3.0","Mode":"global","Name":"traefik_traefik","Ports":"","Replicas":"3/3"}
	{"ID":"vzlr2sm6403z","Image":"ir.goodylabs.com/other/test:6574839","Mode":"replicated","Name":"web_app","Ports":"","Replicas":"3/3 (max 1 per node)"}
	`
	serviceNames, err := services.GetResourcesFromJsonOutput(output)
	assert.NoError(t, err)
	assert.Equal(t, []string{"traefik_traefik", "web_app"}, serviceNames)
}
