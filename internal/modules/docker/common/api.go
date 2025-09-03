package dockercommon

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/modules/docker/common/services"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/config"
)

type EnvCfg struct {
	Name  string
	User  string
	Hosts []string
}

type DockerCommon struct {
	SSHConnector ports.SSHConnector
	Config       map[string]EnvCfg
}

func NewDockerCommon(sshConnector ports.SSHConnector) *DockerCommon {
	return &DockerCommon{
		SSHConnector: sshConnector,
		Config:       make(map[string]EnvCfg),
	}
}
func (d *DockerCommon) LoadConfigFromFile() error {

	baseDir := config.GetBaseDir()
	devopsDirPath := filepath.Join(baseDir, DEVOPS_DIR)

	envs, err := services.ListEnvs(devopsDirPath)
	if err != nil {
		return err
	}

	if len(envs) == 0 {
		return fmt.Errorf("no docker envs found on path %s", devopsDirPath)
	}

	for _, env := range envs {

		scriptPath := filepath.Join(devopsDirPath, env, DEPLOY_FILE)

		var hosts []string
		if targetIp := services.GetSingleIpFromShellFile(scriptPath, TARGET_IP_VAR); targetIp != "" {
			hosts = []string{targetIp}
		}

		if ipAddress := services.GetSingleIpFromShellFile(scriptPath, IP_ADDRESS_VAR); ipAddress != "" {
			hosts = []string{ipAddress}
		}

		if multiIps := services.GetMultipleIpsFromShellScript(scriptPath, IP_ADDRESSES_VAR); len(multiIps) != 0 {
			hosts = multiIps
		}

		if len(hosts) != 0 {
			d.Config[env] = EnvCfg{
				Name:  env,
				User:  "root",
				Hosts: hosts,
			}
		}
	}

	if len(d.Config) == 0 {
		return fmt.Errorf("no valid docker configuration found in %s", devopsDirPath)
	}

	return nil
}

func (d *DockerCommon) GetAvailableEnvs() ([]string, error) {
	if d.Config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	var envs []string
	for env := range d.Config {
		envs = append(envs, env)
	}
	return envs, nil
}

func (d *DockerCommon) GetAvailableHosts(env string) ([]string, error) {
	if d.Config == nil {
		return []string{}, errors.New("Can not get available hosts - config is not loaded")
	}

	return d.Config[env].Hosts, nil
}

func (d *DockerCommon) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	if d.Config == nil {
		return nil, errors.New("Can not get ssh config - config is not loaded")
	}

	return &ports.SSHConfig{
		Host: host,
		User: d.Config[env].User,
		Port: 22,
	}, nil
}
