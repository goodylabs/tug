package app

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/tughelper"
	"github.com/goodylabs/tug/pkg/config"
)

type ConfigureUseCase struct {
	prompter ports.Prompter
}

func NewConfigureUseCase(prompter ports.Prompter) *ConfigureUseCase {
	return &ConfigureUseCase{
		prompter: prompter,
	}
}

func (i *ConfigureUseCase) Execute() error {
	sshDirPath := filepath.Join(config.GetHomeDir(), ".ssh")

	sshFiles, err := tughelper.GetAvailableSSHFiles(sshDirPath)
	if err != nil {
		return err
	}

	sshKeyPath, err := i.prompter.ChooseFromList(sshFiles, "Which SSH key do you want to use?")
	if err != nil {
		return fmt.Errorf("failed to choose SSH key: %w", err)
	}

	tugConfig := tughelper.TugConfig{
		SSHKeyPath: sshKeyPath,
	}
	if err = tughelper.SetTugConfig(&tugConfig); err != nil {
		return err
	}

	fmt.Println("Tug configuration configured successfully!")
	return nil
}
