package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
)

type DeveloperUseCase struct {
	promptAdapter ports.PromptPort
	dockerAdapter ports.DockerPort
}

type DeveloperOptions struct {
	EnvDir string
}

func (d *DeveloperUseCase) Execute(developerOptions *DeveloperOptions) {

	if developerOptions.EnvDir == "" {
		developerOptions.EnvDir = d.getEnvDir()
	}

	targetIp := d.getTargetIp(developerOptions.EnvDir)

	d.dockerAdapter.ConfigureDocker(targetIp)
	containers := d.dockerAdapter.ListContainers()

	chosenContainer := d.choseContainer(containers)
	fmt.Println(chosenContainer)
}

func (d *DeveloperUseCase) getEnvDir() string {
	dirs, _ := adapters.ListDirectories(config.DEVOPS_DIR)
	envDir, _ := d.promptAdapter.ChooseFromList(dirs, "Chose environment")
	return envDir
}

func (d *DeveloperUseCase) getTargetIp(envDir string) string {
	scriptPath := filepath.Join(config.DEVOPS_DIR, envDir, "deploy.sh")
	targetIp, _ := adapters.GetValueFromShFile(scriptPath)
	return targetIp
}

func (d *DeveloperUseCase) choseContainer(containers []dto.ContainerDTO) dto.ContainerDTO {
	containerNames := []string{}
	for _, container := range containers {
		containerNames = append(containerNames, container.Name)
	}
	chosenContainer, _ := d.promptAdapter.ChooseFromList(containerNames, "Chose container")
	for _, container := range containers {
		if container.Name == chosenContainer {
			return container
		}
	}
	panic("Something really went wrong during container chosing process...")
}

func NewDeveloperUseCase() *DeveloperUseCase {
	return &DeveloperUseCase{
		promptAdapter: adapters.NewPromptAdapter(),
		dockerAdapter: adapters.NewDockerAdapter(),
	}
}
