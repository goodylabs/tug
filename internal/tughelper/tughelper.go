package tughelper

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/utils"
)

var (
	tugConfigPath string
)

type TugConfig struct {
	SSHKeyPath string `json:"ssh_key_path"`
}

func GetAvailableSSHFiles(sshDirPath string) ([]string, error) {
	var sshFiles []string
	var err error

	if sshFiles, err = utils.ListFilesInDir(sshDirPath); err != nil {
		return []string{}, fmt.Errorf("Could not list ssh keys on path %s, err: %s", sshDirPath, err)
	}

	for i, file := range sshFiles {
		sshFiles[i] = filepath.Join(sshDirPath, file)
	}

	sshFiles = utils.FilterExclude(sshFiles, ".pub")

	sshFiles = utils.FilterExclude(sshFiles, "known_hosts")

	sshFiles = utils.FilterExclude(sshFiles, "config")

	return sshFiles, nil
}

func GetTugConfig() (*TugConfig, error) {
	var tugConfig *TugConfig
	tugConfigPath := GetTugConfigPath()
	if err := utils.ReadJSON(tugConfigPath, &tugConfig); err != nil {
		return nil, fmt.Errorf("Could not read config from file %s, err: %w", tugConfigPath, err)
	}
	return tugConfig, nil

}

func SetTugConfig(tugConfig *TugConfig) error {
	tugConfigPath := GetTugConfigPath()
	if err := utils.WriteJSON(tugConfigPath, &tugConfig); err != nil {
		return fmt.Errorf("Could not write config to file %s, err: %w", tugConfigPath, err)
	}
	return nil
}

func GetTugConfigPath() string {
	if tugConfigPath == "" {
		tugConfigPath = filepath.Join(config.GetBaseDir(), "tugconfig.json")
	}
	return tugConfigPath
}
