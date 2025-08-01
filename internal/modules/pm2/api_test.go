package pm2_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/modules/pm2"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfigFromFile(t *testing.T) {
	pm2Handler := pm2.NewPm2Handler(nil)

	err := pm2Handler.LoadConfigFromFile()
	assert.NoError(t, err)

	err = pm2Handler.LoadConfigFromFile()
	assert.Error(t, err)
}

func GetAvailableEnvs(t *testing.T) {
	pm2Handler := pm2.NewPm2Handler(nil)
	pm2Handler.LoadConfigFromFile()

	envs, err := pm2Handler.GetAvailableEnvs()
	assert.NoError(t, err)
	assert.Len(t, envs, 6)
}

func TestGetAvailableEnvs(t *testing.T) {
	pm2Handler := pm2.NewPm2Handler(nil)
	pm2Handler.LoadConfigFromFile()

	envs := []string{"staging", "production_RO_2", "production_2"}
	expected := []int{1, 3, 2}

	for i, env := range envs {
		hosts, err := pm2Handler.GetAvailableHosts(env)
		assert.NoError(t, err)
		assert.Len(t, hosts, expected[i])
	}
}

func TestGetSSHConfig(t *testing.T) {
	pm2Handler := pm2.NewPm2Handler(nil)
	pm2Handler.LoadConfigFromFile()

	sshConfig, err := pm2Handler.GetSSHConfig("staging", "xxx.xxx.xxx.xxx")
	assert.NoError(t, err)
	assert.Equal(t, sshConfig.Host, "xxx.xxx.xxx.xxx")
	assert.Equal(t, sshConfig.User, "staging-user")
	assert.Equal(t, sshConfig.Port, 22)

	sshConfig, err = pm2Handler.GetSSHConfig("production_RO_2", "ddd.ddd.ddd.ddd")
	assert.NoError(t, err)
	assert.Equal(t, sshConfig.Host, "ddd.ddd.ddd.ddd")
	assert.Equal(t, sshConfig.User, "root")
	assert.Equal(t, sshConfig.Port, 22)
}

func TestGetAvailableResources(t *testing.T) {

	t.Run("invalid output", func(t *testing.T) {
		pm2Handler := pm2.NewPm2Handler(
			mocks.NewSSHConnectorMock(invalidOutput1, nil),
		)
		pm2Handler.LoadConfigFromFile()

		_, err := pm2Handler.GetAvailableResources(nil)
		assert.ErrorContains(t, err, "failed to parse PM2 list output")
	})

	t.Run("valid output", func(t *testing.T) {
		pm2Handler := pm2.NewPm2Handler(
			mocks.NewSSHConnectorMock(output1, nil),
		)
		pm2Handler.LoadConfigFromFile()

		resources, err := pm2Handler.GetAvailableResources(nil)
		assert.NoError(t, err)
		assert.Len(t, resources, 3)
		assert.Equal(t, resources[0], "pm2-logrotate")
		assert.Equal(t, resources[1], "staging_ro")
		assert.Equal(t, resources[2], "api-staging")
	})

}
