package application

import (
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/tughelper"
)

type InitializeUseCase struct {
	prompter ports.Prompter
}

func NewInitializeUseCase(prompter ports.Prompter) *InitializeUseCase {
	return &InitializeUseCase{
		prompter: prompter,
	}
}

func (i *InitializeUseCase) Execute() error {
	sshDirPath := filepath.Join(config.HOME_DIR, ".ssh")

	sshFiles, err := tughelper.GetAvailableSSHFiles(sshDirPath)
	if err != nil {
		return err
	}

	sshKeyPath, _ := i.prompter.ChooseFromList(sshFiles, "Which SSH key do you want to use?")

	tugConfig := dto.TugConfig{
		SSHKeyPath: sshKeyPath,
	}
	if err = tughelper.SetTugConfig(&tugConfig); err != nil {
		return err
	}
	return nil
}
