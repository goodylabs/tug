package docker_test

// import (
// 	"path/filepath"
// 	"testing"

// 	"github.com/goodylabs/tug/internal/config"
// 	"github.com/goodylabs/tug/internal/constants"
// 	"github.com/goodylabs/tug/internal/services/docker"
// 	"github.com/goodylabs/tug/tests/mocks"
// 	"github.com/stretchr/testify/assert"
// )

// var DockerManager docker.DockerManager

// func init() {
// 	DockerManager = *docker.NewDockerManager(
// 		mocks.NewPrompterMock([]int{}),
// 		mocks.NewSSHConnectorMock("", nil),
// 	)
// }

// func TestGetTargetIpHappyOk(t *testing.T) {
// 	tests := []struct {
// 		envDir      string
// 		resultValue string
// 	}{
// 		{"localhost", "unix:///var/run/docker.sock"},
// 		{"staging", "<ip_address_staging>"},
// 		{"production", "<ip_address_production>"},
// 		{"uat", "<ip_address_uat>"},
// 	}

// 	for _, tt := range tests {
// 		scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, tt.envDir, "deploy.sh")
// 		targetIp, err := DockerManager.GetTargetIP(scriptAbsPath)
// 		assert.Equal(t, tt.resultValue, targetIp)
// 		assert.NoError(t, err, "Expected no error for envDir: %s", tt.envDir)
// 	}
// }

// func TestGetTargetIpNonExistingPath(t *testing.T) {
// 	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, "non-existing-path", "deploy.sh")
// 	targetIp, err := DockerManager.GetTargetIP(scriptAbsPath)
// 	assert.Equal(t, "", targetIp)
// 	assert.Error(t, err)
// }
