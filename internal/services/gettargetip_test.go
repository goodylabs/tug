package services_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/services"
	"github.com/stretchr/testify/assert"
)

func TestGetTargetIpHappyOk(t *testing.T) {
	tests := []struct {
		envDir      string
		resultValue string
	}{
		{"localhost", "unix:///var/run/docker.sock"},
		{"staging", "<ip_address_staging>"},
		{"production", "<ip_address_production>"},
		{"uat", "<ip_address_uat>"},
	}

	for _, tt := range tests {
		scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, tt.envDir, "deploy.sh")
		targetIp, err := services.GetTargetIp(scriptAbsPath)
		assert.Equal(t, tt.resultValue, targetIp)
		assert.NoError(t, err)
	}
}

func TestGetTargetIpNonExistingPath(t *testing.T) {
	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, "non-existing-path", "deploy.sh")

	targetIp, err := services.GetTargetIp(scriptAbsPath)

	assert.Equal(t, "", targetIp)
	assert.Error(t, err)
}
