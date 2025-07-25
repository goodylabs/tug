package application

import (
	"log"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/ports"
)

type DeveloperOptions struct {
	EnvDir string
}

type DockerUseCase struct {
	sshconnector  ports.SSHConnector
	dockerManager ports.DockerManager
}

func NewDockerUseCase(sshAdadpter ports.SSHConnector, dockerManager ports.DockerManager) *DockerUseCase {
	return &DockerUseCase{
		sshconnector:  sshAdadpter,
		dockerManager: dockerManager,
	}
}

func (d *DockerUseCase) Execute(envDir string) {
	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, envDir, "deploy.sh")
	targetIp, err := d.dockerManager.GetTargetIp(scriptAbsPath)
	if err != nil {
		log.Fatal("Error getting target IP:", err)
	}

	err = d.sshconnector.OpenConnection("root", targetIp, 22)
	if err != nil {
		log.Fatal("Error opening SSH connection:", err)
	}
	defer d.sshconnector.CloseConnection()

	containers := d.dockerManager.ListContainers()

	chosenContainer := d.dockerManager.ChoseContainer(containers)

	d.dockerManager.SelectAndExecuteCommand(chosenContainer)
}
