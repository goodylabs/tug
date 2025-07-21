package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
)

type DeveloperUseCase struct {
	prompt_adapter ports.PromptPort
}

func (d *DeveloperUseCase) Execute() {
	dirs, _ := adapters.ListDirectories(config.DEVOPS_DIR)
	fmt.Println(dirs)

	environment, _ := d.prompt_adapter.ChooseFromList(dirs, "Chose environment")
	fmt.Println(environment)

	scriptPath := filepath.Join(config.DEVOPS_DIR, environment, "deploy.sh")
	selectedIp, _ := adapters.GetValueFromShFile(scriptPath)

	fmt.Println(selectedIp)
}

func NewDeveloperUseCase() *DeveloperUseCase {
	return &DeveloperUseCase{
		prompt_adapter: adapters.NewPromptAdapter(),
	}
}
