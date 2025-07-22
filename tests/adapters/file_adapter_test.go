package testadapters

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/constants"

	"github.com/stretchr/testify/assert"
)

func TestListDirectories_ExampleDir(t *testing.T) {
	dirs, err := adapters.ListDirectories(".example-dir")
	assert.NoError(t, err)

	assert.Equal(t, []string{"localhost", "production", "staging"}, dirs)

	log.Println(dirs)
}

func TestGetValueFromShFile(t *testing.T) {
	tests := []struct {
		envDir      string
		resultValue string
	}{
		{"staging", "64.226.87.6"},
		{"production", "167.99.198.9"},
	}

	for _, tt := range tests {
		shFile := filepath.Join(config.BASE_DIR, config.DEVOPS_DIR, tt.envDir, "deploy.sh")
		ipValue, err := adapters.GetValueFromShFile(shFile, constants.TARGET_IP)
		assert.NoError(t, err)
		assert.Equal(t, tt.resultValue, ipValue)
	}
}
