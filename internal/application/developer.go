package application

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/constants"
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
	envDir := developerOptions.EnvDir

	targetIp := d.getTargetIp(envDir)
	stackName := d.getStackName(envDir)

	d.dockerAdapter.ConfigureDocker(targetIp)

	services := d.getServices(stackName)
	fmt.Println(services)
	// containers := d.dockerAdapter.ListServices()

	// chosenContainer := d.choseContainer(containers)
	// fmt.Println(chosenContainer)
}

func (d *DeveloperUseCase) getServices(stackName string) []dto.ServiceDTO {
	services := d.dockerAdapter.ListServices()
	var filtered []dto.ServiceDTO
	for _, service := range services {
		if strings.HasPrefix(service.Name, stackName) {
			filtered = append(filtered, service)
		}
	}
	return filtered
}

func (d *DeveloperUseCase) getEnvDir() string {
	dirs, _ := adapters.ListDirectories(config.DEVOPS_DIR)
	envDir, _ := d.promptAdapter.ChooseFromList(dirs, "Chose environment")
	return envDir
}

func (d *DeveloperUseCase) getTargetIp(envDir string) string {
	scriptPath := filepath.Join(config.DEVOPS_DIR, envDir, "deploy.sh")
	targetIp, _ := adapters.GetValueFromShFile(scriptPath, constants.TARGET_IP)
	return targetIp
}

func (d *DeveloperUseCase) getStackName(envDir string) string {
	scriptPath := filepath.Join(config.DEVOPS_DIR, envDir, "deploy.sh")
	targetIp, _ := adapters.GetValueFromShFile(scriptPath, constants.STACK_NAME)
	return targetIp
}

func (d *DeveloperUseCase) choseContainer(containers []dto.ServiceDTO) dto.ServiceDTO {
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
