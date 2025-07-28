package application

import (
	"log"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/docker"
)

type DeveloperOptions struct {
	EnvDir string
}

type DockerUseCase struct {
	sshconnector  ports.SSHConnector
	DockerManager docker.DockerManager
}

func NewDockerUseCase(sshAdadpter ports.SSHConnector, DockerManager *docker.DockerManager) *DockerUseCase {
	return &DockerUseCase{
		sshconnector:  sshAdadpter,
		DockerManager: *DockerManager,
	}
}

func (d *DockerUseCase) Execute(envDir string) {
	var targetIp string
	var err error

	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, envDir, "deploy.sh")

	if targetIp, err = d.DockerManager.GetTargetIp(scriptAbsPath); err != nil {
		log.Fatal("Error getting target IP:", err)
	}

	if err = d.sshconnector.OpenConnection("root", targetIp, 22); err != nil {
		log.Fatal("Error opening SSH connection:", err)
	}
	defer d.sshconnector.CloseConnection()

	containers := d.DockerManager.ListContainers()

	selectedContainer := d.DockerManager.SelectContainer(containers)

	d.DockerManager.SelectAndExecuteCommand(selectedContainer)
}
