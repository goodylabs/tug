package loadproject

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/modules"
	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/utils"
)

const (
	DEVOPS_DIR       = "devops"
	DEPLOY_FILE      = "deploy.sh"
	TARGET_IP_VAR    = "TARGET_IP"
	IP_ADDRESS_VAR   = "IP_ADDRESS"
	IP_ADDRESSES_VAR = "IP_ADDRESSES"
)

type DockerLoadStrategy struct{}

func NewDockerLoadStrategy() *DockerLoadStrategy {
	return &DockerLoadStrategy{}
}

func (s *DockerLoadStrategy) Execute() (modules.ProjectConfig, error) {
	projectCfg := modules.ProjectConfig{
		Config: make(map[string]modules.EnvConfig),
	}

	baseDir := config.GetBaseDir()
	devopsDirPath := filepath.Join(baseDir, DEVOPS_DIR)

	envs, err := utils.ListDirsOnPath(devopsDirPath)
	if err != nil {
		return projectCfg, fmt.Errorf("docker strategy: cannot read dirs: %w", err)
	}

	for _, env := range envs {
		scriptPath := filepath.Join(devopsDirPath, env, DEPLOY_FILE)
		hosts := s.parseHosts(scriptPath)

		if len(hosts) > 0 {
			projectCfg.Config[env] = modules.EnvConfig{
				Name:  env,
				User:  "root",
				Hosts: hosts,
			}
		}
	}

	if len(projectCfg.Config) == 0 {
		return projectCfg, fmt.Errorf("no valid docker configuration in %s", devopsDirPath)
	}

	return projectCfg, nil
}

func (s *DockerLoadStrategy) parseHosts(path string) []string {
	if hosts := s.getMultipleIps(path, IP_ADDRESSES_VAR); len(hosts) > 0 {
		return hosts
	}
	if host := s.getSingleVar(path, IP_ADDRESS_VAR); host != "" {
		return []string{host}
	}
	if host := s.getSingleVar(path, TARGET_IP_VAR); host != "" {
		return []string{host}
	}
	return nil
}

func (s *DockerLoadStrategy) getSingleVar(path, key string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if after, ok := strings.CutPrefix(line, key+"="); ok {
			val := after
			return strings.Trim(val, `"' `+"`")
		}
	}
	return ""
}

func (s *DockerLoadStrategy) getMultipleIps(path, key string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if !strings.HasPrefix(line, key+"=") {
			continue
		}

		raw := strings.TrimPrefix(line, key+"=")
		cleanRaw := strings.Trim(raw, `()"' `+"`")

		parts := strings.Fields(cleanRaw)

		var result []string
		for _, part := range parts {
			cleanPart := strings.Trim(part, `"'`+"`")
			if cleanPart != "" {
				result = append(result, cleanPart)
			}
		}
		return result
	}
	return nil
}
