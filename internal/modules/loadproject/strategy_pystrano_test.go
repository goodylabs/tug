package loadproject_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/modules/loadproject"
	"github.com/stretchr/testify/assert"
)

func TestPystranoLoadStrategy_Execute(t *testing.T) {
	lp := loadproject.NewLoadProject()

	t.Run("successful load pystrano environments", func(t *testing.T) {
		cfg, err := lp.Execute(loadproject.PystranoStrategy)

		assert.NoError(t, err)
		assert.NotEmpty(t, cfg.Config)

		cmsProdKey := filepath.Join("cms", "production")
		if env, ok := cfg.Config[cmsProdKey]; ok {
			assert.Len(t, env.Hosts, 4)
			assert.Equal(t, "c_p_ipv4_nr_2", env.Hosts[1])
		} else {
			t.Errorf("Environment %s not found in config", cmsProdKey)
		}

		schedulerStagingKey := filepath.Join("scheduler", "staging")
		if env, ok := cfg.Config[schedulerStagingKey]; ok {
			assert.Len(t, env.Hosts, 1)
			assert.Equal(t, "s_s_ipv4_nr_1", env.Hosts[0])
		}
	})
}
