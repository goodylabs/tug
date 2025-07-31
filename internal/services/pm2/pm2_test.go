package pm2_test

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

const output1 = `[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]`

const invalidOutput1 = `>>>> In-memory PM2 is out-of-date, do:
>>>> $ pm2 update
In memory PM2 version: 5.2.2
Local PM2 version: 6.0.8

[{ "name": "pm2-logrotate", "pid": 3105},{  "name": "staging_ro", "pid": 4043331},{ "name": "api-staging", "pid": 4041884}]`

func init() {
	config.Load()
}

func TestRetrievePm2Config(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	config.Load()
	fmt.Println("BASE_DIR:", config.BASE_DIR)

	pm2ConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	pm2Config, err := pm2Manager.RetrievePm2Config(pm2ConfigPath)

	assert.NoError(t, err)
	assert.NotNil(t, pm2Config)

	assert.Equal(t, pm2Config.Apps[0].Name, "pm2-app-1")
	assert.Equal(t, pm2Config.Deploy["staging"].User, "staging-user")
	assert.Equal(t, pm2Config.Deploy["staging"].Host[0], "xxx.xxx.xxx.xxx")

	envs := pm2Config.ListEnvironments()
	requiredEnvs := []string{"staging", "staging_RO", "production_1", "production_2", "production_RO_1", "production_RO_2"}
	assert.Len(t, envs, 6)
	for _, env := range envs {
		assert.Contains(t, requiredEnvs, env)
	}
}

// func TestRetrievePm2ConfigInvalidFile(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

// 	_, err := pm2Manager.RetrievePm2Config(filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE))
// 	assert.ErrorContains(t, err, "cannot read json file")
// }

// func TestRetrievePm2ConfigNodeScriptFails(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

// 	_, err := pm2Manager.RetrievePm2Config(filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE))
// 	assert.ErrorContains(t, err, "Can not load config from file(probably doesn't")
// }

// func TestGetAvailableEnvsBadArg(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

// 	envs, err := pm2Manager.GetAvailableEnvs()
// 	assert.ErrorContains(t, err, "no environments found in PM2 config")
// 	assert.True(t, len(envs) == 0)
// }

// func TestGetAvailableEnvsEmptyArg(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, "", nil)

// 	env, err := pm2Manager.GetAvailableEnvs()
// 	assert.NoError(t, err)
// 	assert.Equal(t, "production_1", env)
// }

// func TestGetAvailableEnvsOkArg(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

// 	env, err := pm2Manager.GetAvailableEnvs()
// 	assert.NoError(t, err)
// 	assert.Equal(t, "staging_RO", env)
// }

func TestGetSSHConfigAutoSelectHost(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{}, "", nil)

	sshConfig, err := pm2Manager.GetSSHConfig("staging_RO", "yyy.yyy.yyy.yyy")

	assert.NoError(t, err)
	assert.Equal(t, sshConfig.User, "staging-user")
	assert.Equal(t, sshConfig.Host, "yyy.yyy.yyy.yyy")
	assert.Equal(t, sshConfig.Port, 22)
}

func TestGetSSHConfigSelectSecondHost(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{1}, "", nil)

	sshConfig, err := pm2Manager.GetSSHConfig("production_RO_2", "ddd.ddd.ddd.ddd")

	assert.NoError(t, err)
	assert.Equal(t, sshConfig.User, "root")
	assert.Equal(t, sshConfig.Host, "ddd.ddd.ddd.ddd")
	assert.Equal(t, sshConfig.Port, 22)
}

// func TestGetSSHConfigDummyEnv(t *testing.T) {
// 	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{1}, "", nil)

// 	sshConfig, err := pm2Manager.GetSSHConfig("dummy_env")

// 	assert.ErrorContains(t, err, "environment 'dummy_env' not found in loaded PM2 config")
// 	assert.Nil(t, sshConfig)
// }
