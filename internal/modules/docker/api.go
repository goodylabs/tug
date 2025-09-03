package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/modules/docker/services"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/config"
)

const (
	devopsDir = "devops"
)

type envCfg struct {
	Name  string
	User  string
	Hosts []string
}

type DockerManager struct {
	sshConnector ports.SSHConnector
	config       *map[string]envCfg
}

func NewDockerManager(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &DockerManager{
		sshConnector: sshConnector,
	}
}

func (d *DockerManager) LoadConfigFromFile() error {

	baseDir := config.GetBaseDir()
	devopsDirPath := filepath.Join(baseDir, devopsDir)

	envs, err := services.ListEnvs(devopsDirPath)
	if err != nil {
		return err
	}

	for _, env := range envs {

		scriptPath := filepath.Join(devopsDirPath, env, "deploy.sh")

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

		if len(hosts) == 0 {
			(*d.config)[env] = envCfg{
				Name:  env,
				User:  "root",
				Hosts: hosts,
			}
		}
	}

	if len(*d.config) == 0 {
		return fmt.Errorf("no valid docker configuration found in %s", devopsDirPath)
	}

	return nil
}

func (d *DockerManager) GetAvailableEnvs() ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	envs := make([]string, 0, len((*d.config)))
	for env := range *d.config {
		envs = append(envs, env)
	}
	return envs, nil
}

func (d *DockerManager) GetAvailableHosts(env string) ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available hosts - config is not loaded")
	}

	return (*d.config)[env].Hosts, nil
}

func (d *DockerManager) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	return &ports.SSHConfig{
		Host: host,
		User: (*d.config)[env].User,
		Port: 22,
	}, nil
}

func (d *DockerManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	var containers []containerDTO

	output, err := d.sshConnector.RunCommand(dockerListCmd)
	if err != nil {
		return nil, err
	}

	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")
	for line := range lines {
		var container containerDTO
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			continue
		}
		containers = append(containers, container)
	}

	containerNames := make([]string, len(containers))
	for i, c := range containers {
		containerNames[i] = c.Name
	}

	return containerNames, nil
}
func (d *DockerManager) GetAvailableActionTemplates() map[string]string {
	return commandTemplates
}
