package loadproject_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/stretchr/testify/assert"
)

func TestDockerLoadStrategy_Legacy(t *testing.T) {
	strategy := loadproject.NewDockerLoadStrategy()

	t.Run("Check ListEnvs logic via Execute", func(t *testing.T) {
		cfg, err := strategy.Execute()

		assert.NoError(t, err)
		assert.Contains(t, cfg.Config, "localhost")
		assert.Contains(t, cfg.Config, "production")
	})

	t.Run("IP_ADDRESS variable logic", func(t *testing.T) {
		cfg, err := strategy.Execute()
		assert.NoError(t, err)

		expectedHost := "unix:///var/run/docker.sock"
		assert.Equal(t, []string{expectedHost}, cfg.Config["localhost"].Hosts)
	})

	t.Run("TARGET_IP variable logic", func(t *testing.T) {
		cfg, err := strategy.Execute()
		assert.NoError(t, err)

		expectedHost := "<ip_address_production>"
		assert.Equal(t, []string{expectedHost}, cfg.Config["production"].Hosts)
	})

	t.Run("Multiple IPs logic", func(t *testing.T) {
		cfg, err := strategy.Execute()
		assert.NoError(t, err)

		expectedHosts := []string{"ip_1", "ip_2", "ip_3"}
		assert.Equal(t, expectedHosts, cfg.Config["production_v2"].Hosts)
	})
}

func TestLoadProject_Selector(t *testing.T) {
	lp := loadproject.NewLoadProject()

	t.Run("Execute with Docker Strategy", func(t *testing.T) {
		cfg, err := lp.Execute(loadproject.DockerStrategy)
		assert.NoError(t, err)
		assert.NotEmpty(t, cfg.Config)
	})

	t.Run("Execute with unsupported strategy", func(t *testing.T) {
		_, err := lp.Execute("non-existent")
		assert.ErrorContains(t, err, "unsupported strategy")
	})
}
