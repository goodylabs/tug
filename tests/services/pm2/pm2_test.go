package pm2_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoadPm2ConfigOk(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig
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

func TestLoadPm2ConfigInvalidFile(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, "ecosystem.config.invalid.js")
	err := pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	assert.ErrorContains(t, err, "cannot read json file", "Expected error when loading invalid ecosystem config")
}

func TestLoadPm2ConfigNodeScriptFails(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, "ecosystem.config.nonexisting.js")
	err := pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)
	assert.ErrorContains(t, err, "Can not load config from file(probably doesn't")
}

func TestSelectEnvFromConfigBadArg(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, "ecosystem.config.emptyhosts.js")
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	env, err := pm2Manager.SelectEnvFromConfig(&pm2Config, "")
	assert.ErrorContains(t, err, "no environments found in PM2 config")
	assert.Equal(t, "", env)
}

func TestSelectEnvFromConfigEmptyArg(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	env, err := pm2Manager.SelectEnvFromConfig(&pm2Config, "")
	assert.NoError(t, err, "Expected no error when selecting environment from config")
	assert.Equal(t, "production_1", env, "Expected environment to be 'production_1' when no argument is provided")
}

func TestSelectEnvFromConfigOkArg(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	env, err := pm2Manager.SelectEnvFromConfig(&pm2Config, "staging_RO")
	assert.NoError(t, err, "Expected no error when selecting environment from config")
	assert.Equal(t, "staging_RO", env, "Expected environment to be 'staging_RO' when valid argument is provided")
}

func TestGetSSHConfigAutoSelectHost(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	sshConfig, err := pm2Manager.GetSSHConfig(&pm2Config, "staging_RO")

	assert.NoError(t, err, "Expected no error when getting SSH config")
	assert.Equal(t, sshConfig.User, "staging-user")
	assert.Equal(t, sshConfig.Host, "yyy.yyy.yyy.yyy")
	assert.Equal(t, sshConfig.Port, 22)
}

func TestGetSSHConfigSelectSecondHost(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{1}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)
	sshConfig, err := pm2Manager.GetSSHConfig(&pm2Config, "production_RO_2")

	assert.NoError(t, err, "Expected no error when getting SSH config")
	assert.Equal(t, sshConfig.User, "root")
	assert.Equal(t, sshConfig.Host, "ddd.ddd.ddd.ddd")
	assert.Equal(t, sshConfig.Port, 22)
}

func TestGetSSHConfigDummyEnv(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{1}, "", nil)

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config)

	sshConfig, err := pm2Manager.GetSSHConfig(&pm2Config, "dummy_env")

	assert.ErrorContains(t, err, "environment 'dummy_env' not found in loaded PM2 config")
	assert.Nil(t, sshConfig)
}
