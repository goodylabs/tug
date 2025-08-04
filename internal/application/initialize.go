package application

import (
	"fmt"
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

	sshKeyPath, err := i.prompter.ChooseFromList(sshFiles, "Which SSH key do you want to use?")
	if err != nil {
		return fmt.Errorf("failed to choose SSH key: %w", err)
	}

	tugConfig := dto.TugConfig{
		SSHKeyPath: sshKeyPath,
	}
	if err = tughelper.SetTugConfig(&tugConfig); err != nil {
		return err
	}

	fmt.Println("Tug configuration initialized successfully!")
	return nil
}
