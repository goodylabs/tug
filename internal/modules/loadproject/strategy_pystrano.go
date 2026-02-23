package loadproject

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/utils"
)

type PystranoLoadStrategy struct{}

func NewPystranoLoadStrategy() *PystranoLoadStrategy {
	return &PystranoLoadStrategy{}
}

func (s *PystranoLoadStrategy) Execute() (ProjectConfig, error) {
	projectCfg := ProjectConfig{
		Config: make(map[string]EnvConfig),
	}

	pystranoDir := filepath.Join(config.GetBaseDir(), "deploy")

	deploymentFiles, err := s.findDeploymentYAML(pystranoDir)
	if err != nil {
		return projectCfg, fmt.Errorf("pystrano strategy: %w", err)
	}

	if len(deploymentFiles) == 0 {
		return projectCfg, fmt.Errorf("no pystrano config files found in 'deploy' directory")
	}

	sort.Strings(deploymentFiles)

	for _, relPath := range deploymentFiles {
		fullPath := filepath.Join(pystranoDir, relPath)

		hosts, err := s.retrieveHosts(fullPath)
		if err != nil {
			continue
		}

		if len(hosts) > 0 {
			envName := filepath.Dir(relPath)

			projectCfg.Config[envName] = EnvConfig{
				Name:  envName,
				User:  "root",
				Hosts: hosts,
			}
		}
	}

	return projectCfg, nil
}

func (s *PystranoLoadStrategy) findDeploymentYAML(baseDir string) ([]string, error) {
	var result []string

	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			base := filepath.Base(path)
			if base == "deployment.yml" || base == "deployment.yaml" {
				rel, err := filepath.Rel(baseDir, path)
				if err != nil {
					return err
				}
				result = append(result, rel)
			}
		}
		return nil
	})

	return result, err
}

type pystranoConfigFile struct {
	Servers []struct {
		Host string `yaml:"host"`
	} `yaml:"servers"`
}

func (s *PystranoLoadStrategy) retrieveHosts(filePath string) ([]string, error) {
	var cfg pystranoConfigFile
	if err := utils.ReadYAML(filePath, &cfg); err != nil {
		return nil, err
	}

	var hosts []string
	for _, server := range cfg.Servers {
		if server.Host != "" {
			hosts = append(hosts, server.Host)
		}
	}
	return hosts, nil
}
