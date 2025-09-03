package docker

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/modules/docker/services"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/config"
)

type envCfg struct {
	Name  string
	User  string
	Hosts []string
}

type DockerManager struct {
	sshConnector ports.SSHConnector
	config       map[string]envCfg
}

func NewDockerManager(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &DockerManager{
		sshConnector: sshConnector,
		config:       make(map[string]envCfg),
	}
}

func (d *DockerManager) LoadConfigFromFile() error {

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
			d.config[env] = envCfg{
				Name:  env,
				User:  "root",
				Hosts: hosts,
			}
		}
	}

	if len(d.config) == 0 {
		return fmt.Errorf("no valid docker configuration found in %s", devopsDirPath)
	}

	return nil
}

func (d *DockerManager) GetAvailableEnvs() ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	var envs []string
	for env := range d.config {
		envs = append(envs, env)
	}
	return envs, nil
}

func (d *DockerManager) GetAvailableHosts(env string) ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available hosts - config is not loaded")
	}

	return d.config[env].Hosts, nil
}

func (d *DockerManager) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	if d.config == nil {
		return nil, errors.New("Can not get ssh config - config is not loaded")
	}

	return &ports.SSHConfig{
		Host: host,
		User: d.config[env].User,
		Port: 22,
	}, nil
}

func (d *DockerManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available resources - config is not loaded")
	}

	dockerListCmd := "docker ps --format json"
	output, err := d.sshConnector.RunCommand(dockerListCmd)
	if err != nil {
		return nil, err
	}

	return services.GetResourcesFromJsonOutput(output)
}

func (d *DockerManager) GetAvailableActionTemplates() map[string]string {
	return services.GetActionTemplates()
}
