package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/utils"
)

type DockerManager struct {
	sshConnector ports.SSHConnector
	config       *dockerConfig
}

func NewDockerManager(prompter ports.Prompter, sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &DockerManager{
		sshConnector: sshConnector,
	}
}

func (d *DockerManager) LoadConfigFromFile() error {
	devopsDirPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR)
	environments, err := utils.ListDirsOnPath(devopsDirPath)
	if err != nil {
		return err
	}

	d.config = &dockerConfig{
		Envs: make(map[string]dockerConfigEnv),
	}

	var path string
	for _, env := range environments {
		path = filepath.Join(devopsDirPath, env, constants.DOCKER_CONFIG_FILE)
		if _, err := os.Stat(path); err != nil {
			continue
		}

		targetIp, err := d.GetTargetIpFromScript(filepath.Join(devopsDirPath, env, constants.DOCKER_CONFIG_FILE))
		if err != nil {
			continue
		}

		d.config.Envs[env] = dockerConfigEnv{
			Name:  env,
			User:  "root",
			Hosts: []string{targetIp},
		}
	}

	if len(d.config.Envs) == 0 {
		return fmt.Errorf("no valid docker configuration found in %s", devopsDirPath)
	}

	return nil
}
func (d *DockerManager) GetAvailableEnvs() ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	envs := make([]string, 0, len(d.config.Envs))
	for env := range d.config.Envs {
		envs = append(envs, env)
	}
	return envs, nil
}
func (d *DockerManager) GetAvailableHosts(env string) ([]string, error) {
	if d.config == nil {
		return []string{}, errors.New("Can not get available hosts - config is not loaded")
	}

	return d.config.Envs[env].Hosts, nil
}
func (d *DockerManager) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	return &ports.SSHConfig{
		Host: host,
		User: d.config.Envs[env].User,
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
