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
		return fmt.Errorf("getting SSH directory path: %w", err)
	}

	sshKeyPath := i.prompter.ChooseFromList(sshFiles, "Which SSH key do you want to use?")

	tugConfig := dto.TugConfig{
		SSHKeyPath: sshKeyPath,
	}
	if err = tughelper.SetTugConfig(&tugConfig); err != nil {
		return fmt.Errorf("setting Tug config: %w", err)
	}

	fmt.Println("Tug configuration initialized successfully.")
	return nil
}
