package tughelper

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

type TugHelper struct {
	prompter ports.Prompter
}

func NewTugHelper(prompter ports.Prompter) *TugHelper {
	return &TugHelper{
		prompter: prompter,
	}
}

func (t *TugHelper) GetSSHDirPath(sshDirPath string) (string, error) {
	var sshFiles []string
	var err error

	if sshFiles, err = utils.ListFilesInDir(sshDirPath); err != nil {
		return "", fmt.Errorf("Occurred error while listing files in %s directory: %w", sshDirPath, err)
	}

	for i, file := range sshFiles {
		sshFiles[i] = filepath.Join(sshDirPath, file)
	}

	sshFiles = utils.FilterExclude(sshFiles, ".pub")

	sshFiles = utils.FilterExclude(sshFiles, "known_hosts")

	sshFiles = utils.FilterExclude(sshFiles, "config")

	return t.prompter.ChooseFromList(sshFiles, "Which ssh key u want to use to connect with servers?"), nil
}
