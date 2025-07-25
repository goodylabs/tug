package application

import (
	"log"
	"path/filepath"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/services"
)

type DeveloperOptions struct {
	EnvDir string
}

func DeveloperUseCase(devOptions *DeveloperOptions) {
	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, devOptions.EnvDir, "deploy.sh")
	targetIp, err := services.GetTargetIp(scriptAbsPath)

	log.Printf("Connecting to %s...", targetIp)

	if err != nil {
		log.Fatal(err)
	}

	adapters.DockerApi.ConfigureDocker(targetIp)

	containers := adapters.DockerApi.ListContainers()

	chosenContainer := services.ChoseContainer(containers)

	services.SelectAndExecCommandOnDockerContainer(chosenContainer)
}
