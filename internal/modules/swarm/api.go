package swarm

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

type SwarmManager struct {
	sshConnector ports.SSHConnector
	config       *dockerConfig
}

func NewSwarmManager(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &SwarmManager{
		sshConnector: sshConnector,
	}
}

func (s *SwarmManager) LoadConfigFromFile() error {
	devopsDirPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR)
	environments, err := utils.ListDirsOnPath(devopsDirPath)
	if err != nil {
		return err
	}

	s.config = &dockerConfig{
		Envs: make(map[string]dockerConfigEnv),
	}

	var path string
	for _, env := range environments {
		path = filepath.Join(devopsDirPath, env, constants.DOCKER_CONFIG_FILE)
		if _, err := os.Stat(path); err != nil {
			continue
		}

		targetIp, err := s.GetTargetIpFromScript(filepath.Join(devopsDirPath, env, constants.DOCKER_CONFIG_FILE))
		if err != nil {
			continue
		}

		s.config.Envs[env] = dockerConfigEnv{
			Name:  env,
			User:  "root",
			Hosts: []string{targetIp},
		}
	}

	if len(s.config.Envs) == 0 {
		return fmt.Errorf("no valid docker configuration found in %s", devopsDirPath)
	}

	return nil
}
func (s *SwarmManager) GetAvailableEnvs() ([]string, error) {
	if s.config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	envs := make([]string, 0, len(s.config.Envs))
	for env := range s.config.Envs {
		envs = append(envs, env)
	}
	return envs, nil
}
func (s *SwarmManager) GetAvailableHosts(env string) ([]string, error) {
	if s.config == nil {
		return []string{}, errors.New("Can not get available hosts - config is not loaded")
	}

	return s.config.Envs[env].Hosts, nil
}
func (s *SwarmManager) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	return &ports.SSHConfig{
		Host: host,
		User: s.config.Envs[env].User,
		Port: 22,
	}, nil
}

func (s *SwarmManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	var containers []serviceDTO

	output, err := s.sshConnector.RunCommand(swarmListCmd)
	if err != nil {
		return nil, err
	}

	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")
	for line := range lines {
		var container serviceDTO
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
func (s *SwarmManager) GetAvailableActionTemplates() map[string]string {
	return commandTemplates
}
