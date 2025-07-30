package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/tughelper"
)

type InitializeUseCase struct {
	prompter  ports.Prompter
	tugHelper *tughelper.TugHelper
}

func NewInitializeUseCase(prompter ports.Prompter, tughelper *tughelper.TugHelper) *InitializeUseCase {
	return &InitializeUseCase{
		prompter:  prompter,
		tugHelper: tughelper,
	}
}

func (i *InitializeUseCase) Execute() error {

	sshDirPath := filepath.Join(config.HOME_DIR, ".ssh")

	sshKeyPath, err := i.tugHelper.GetSSHDirPath(sshDirPath)
	fmt.Println(sshKeyPath)
	if err != nil {
		return fmt.Errorf("getting SSH directory path: %w", err)
	}

	return nil
}
