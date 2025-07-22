package application

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/goodylabs/docker-swarm-cli/internal/adapters"
	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/constants"
	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
	"github.com/goodylabs/docker-swarm-cli/internal/services"
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

	d.dockerAdapter.ConfigureDocker(targetIp)

	containers := d.dockerAdapter.ListContainers()

	chosenContainer := d.choseContainer(containers)

	services.SelectCommandsService(d.promptAdapter, chosenContainer)
}

func (d *DeveloperUseCase) choseContainer(containers []dto.ContainerDTO) dto.ContainerDTO {
	var names []string
	for _, container := range containers {
		names = append(names, container.Name)
	}
	selectedName := d.promptAdapter.ChooseFromList(names, "Chose container")
	for _, container := range containers {
		if selectedName == container.Name {
			return container
		}
	}
	panic(constants.PANIC)
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

// func (d *DeveloperUseCase) getEnvDir() string {
// 	dirs, _ := adapters.ListDirectories(config.DEVOPS_DIR)
// 	envDir, _ := d.promptAdapter.ChooseFromList(dirs, "Chose environment")
// 	return envDir
// }

func (d *DeveloperUseCase) getTargetIp(envDir string) string {

	scriptPath := filepath.Join(config.DEVOPS_DIR, envDir, "deploy.sh")

	legacyTargetIp, _ := adapters.GetValueFromShFile(scriptPath, constants.LEGACY_TARGET_IP_KEY)
	if legacyTargetIp != "" {
		return legacyTargetIp
	}

	targetIp, err := adapters.GetValueFromShFile(scriptPath, constants.TARGET_IP_KEY)
	if err != nil {
		log.Fatal(err)
	}

	return targetIp
}

func (d *DeveloperUseCase) getStackName(envDir string) string {
	scriptPath := filepath.Join(config.DEVOPS_DIR, envDir, "deploy.sh")
	targetIp, _ := adapters.GetValueFromShFile(scriptPath, constants.STACK_NAME)
	return targetIp
}

func NewDeveloperUseCase() *DeveloperUseCase {
	return &DeveloperUseCase{
		promptAdapter: adapters.NewPromptAdapter(),
		dockerAdapter: adapters.NewDockerAdapter(),
	}
}
