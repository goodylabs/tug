package services_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/modules/docker"
	"github.com/goodylabs/tug/internal/modules/docker/services"
	"github.com/goodylabs/tug/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestListEnvs(t *testing.T) {

	var devopsDirPath = filepath.Join(config.GetBaseDir(), "devops")

	t.Run("existing path", func(t *testing.T) {
		envs, err := services.ListEnvs(devopsDirPath)

		assert.NoError(t, err)
		assert.Equal(t, []string{"localhost", "production"}, envs[0:2])
	})

	t.Run("non existing path", func(t *testing.T) {
		_, err := services.ListEnvs("404")

		assert.ErrorContains(t, err, "open 404: no such file")
	})
}

func TestGetSingleIpFromShellFile(t *testing.T) {

	var devopsDirPath = filepath.Join(config.GetBaseDir(), "devops")

	t.Run("IP_ADDRESS variable", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "localhost", "deploy.sh")

		ip := services.GetSingleIpFromShellFile(scriptPath, docker.IP_ADDRESS_VAR)

		assert.Equal(t, "unix:///var/run/docker.sock", ip)
	})

	t.Run("TARGET_IP variable", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "production", "deploy.sh")

		ip := services.GetSingleIpFromShellFile(scriptPath, docker.TARGET_IP_VAR)

		assert.Equal(t, "<ip_address_production>", ip)
	})

	t.Run("empty line on non-existing variable", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "production", "deploy.sh")

		ip := services.GetSingleIpFromShellFile(scriptPath, "404")

		assert.Equal(t, "", ip)
	})
}

func TestGetMultipleIpsFromShellFile(t *testing.T) {

	var devopsDirPath = filepath.Join(config.GetBaseDir(), "devops")

	t.Run("IP_ADDRESS variable", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "production_v2", "deploy.sh")

		ip := services.GetMultipleIpsFromShellScript(scriptPath, docker.IP_ADDRESSES_VAR)

		assert.Equal(t, []string{"ip_1", "ip_2", "ip_3"}, ip)
	})

	t.Run("empty line on non-existing variable", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "production", "404")

		ip := services.GetMultipleIpsFromShellScript(scriptPath, "")

		assert.Len(t, ip, 0)
	})

	t.Run("empty line on non-existing file", func(t *testing.T) {
		scriptPath := filepath.Join(devopsDirPath, "production", "deploy.sh")

		ip := services.GetMultipleIpsFromShellScript(scriptPath, "404")

		assert.Len(t, ip, 0)
	})
}
