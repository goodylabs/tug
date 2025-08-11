package tughelper

import (
	"path/filepath"

	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/pkg/utils"
)

func GetAvailableSSHFiles(sshDirPath string) ([]string, error) {
	var sshFiles []string
	var err error

	if sshFiles, err = utils.ListFilesInDir(sshDirPath); err != nil {
		return []string{}, err
	}

	for i, file := range sshFiles {
		sshFiles[i] = filepath.Join(sshDirPath, file)
	}

	sshFiles = utils.FilterExclude(sshFiles, ".pub")

	sshFiles = utils.FilterExclude(sshFiles, "known_hosts")

	sshFiles = utils.FilterExclude(sshFiles, "config")

	return sshFiles, err
}

func GetTugConfig() (*TugConfig, error) {
	var tugConfig *TugConfig
	err := utils.ReadJSON(config.TUG_CONFIG_PATH, &tugConfig)
	return tugConfig, err
}

func SetTugConfig(tugConfig *TugConfig) error {
	return utils.WriteJSON(config.TUG_CONFIG_PATH, &tugConfig)
}
