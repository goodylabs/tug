package services

import (
	"os"
	"path/filepath"

	"github.com/goodylabs/tug/pkg/utils"
)

func FindDeploymentYAML(baseDir string) ([]string, error) {
	var result []string

	err := filepath.WalkDir(baseDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && (filepath.Base(path) == "deployment.yml" || filepath.Base(path) == "deployment.yaml") {
			rel, err := filepath.Rel(baseDir, path)
			if err != nil {
				return err
			}
			result = append(result, rel)
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

func RetrieveHostsFromConfigFile(filePath string) ([]string, error) {
	var cfg pystranoConfigFile
	if err := utils.ReadYAML(filePath, &cfg); err != nil {
		return nil, err
	}
	var hosts []string
	for _, server := range cfg.Servers {
		hosts = append(hosts, server.Host)
	}
	return hosts, nil
}
