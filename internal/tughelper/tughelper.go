package tughelper

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/utils"
)

func GetAvailableSSHFiles(sshDirPath string) ([]string, error) {
	var sshFiles []string
	var err error

	if sshFiles, err = utils.ListFilesInDir(sshDirPath); err != nil {
		return []string{}, fmt.Errorf("Occurred error while listing files in %s directory: %w", sshDirPath, err)
	}

	for i, file := range sshFiles {
		sshFiles[i] = filepath.Join(sshDirPath, file)
	}

	sshFiles = utils.FilterExclude(sshFiles, ".pub")

	sshFiles = utils.FilterExclude(sshFiles, "known_hosts")

	sshFiles = utils.FilterExclude(sshFiles, "config")

	return sshFiles, err
}

func GetTugConfig() (*dto.TugConfig, error) {
	var tugConfig *dto.TugConfig
	err := utils.ReadJSON(config.TUG_CONFIG_PATH, &tugConfig)
	return tugConfig, err
}

func SetTugConfig(tugConfig *dto.TugConfig) error {
	return utils.WriteJSON(config.TUG_CONFIG_PATH, &tugConfig)
}
