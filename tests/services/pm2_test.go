package services_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/services/pm2"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

var pm2Manager *pm2.Pm2Manager

func init() {
	pm2Manager = pm2.NewPm2Manager(
		mocks.NewPrompterMock([]int{}),
		mocks.NewSSHConnectorMock(),
	)
}

func TestGetEcosystemConfig(t *testing.T) {
	var pm2Config dto.EconsystemConfigDTO

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	err := pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	assert.NoError(t, err, "Expected no error when loading ecosystem config")
	assert.NotNil(t, pm2Config, "Expected pm2Config to be not nil")

	assert.Equal(t, pm2Config.Apps[0].Name, "pm2-app-1", "Expected app name to be 'pm2-app-1'")
	assert.Equal(t, pm2Config.Deploy["staging"].User, "staging-user", "Expected staging user to be 'staging-user'")
	assert.Equal(t, pm2Config.Deploy["staging"].Host[0], "xxx.xxx.xxx.xxx", "Expected staging host to be 'xxx.xxx.xxx.xxx'")

	envs := pm2Config.ListEnvironments()
	requiredEnvs := []string{"staging", "staging_RO", "production_1", "production_2", "production_RO_1", "production_RO_2"}
	assert.Len(t, envs, 6, "Expected 6 environments in the config")
	for _, env := range envs {
		assert.Contains(t, requiredEnvs, env, "Expected environment to be either 'staging' or 'production'")
	}
}
