package services_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/modules/pystrano/services"
	"github.com/goodylabs/tug/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestFindDeploymentYAML(t *testing.T) {
	pystranoDir := filepath.Join(config.GetBaseDir(), "deploy")
	deployFiles, err := services.FindDeploymentYAML(pystranoDir)
	assert.NoError(t, err)
	assert.Len(t, deployFiles, 4)
	assert.Equal(t, deployFiles[0], filepath.Join("cms", "production", "deployment.yml"))
}

func TestRetrieveHostsFromConfigFile(t *testing.T) {
	pystranoDir := filepath.Join(config.GetBaseDir(), "deploy")

	t.Run("cms production", func(t *testing.T) {
		cmsProdFile := filepath.Join(pystranoDir, "cms", "production", "deployment.yml")
		hosts, err := services.RetrieveHostsFromConfigFile(cmsProdFile)
		assert.NoError(t, err)
		assert.Len(t, hosts, 4)
		assert.Equal(t, hosts[1], "c_p_ipv4_nr_2")
	})

	t.Run("scheduler staging", func(t *testing.T) {
		schedulerStagingFile := filepath.Join(pystranoDir, "scheduler", "staging", "deployment.yml")
		hosts, err := services.RetrieveHostsFromConfigFile(schedulerStagingFile)
		assert.NoError(t, err)
		assert.Len(t, hosts, 1)
		assert.Equal(t, hosts[0], "s_s_ipv4_nr_1")
	})
}
