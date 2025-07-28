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
	sshConnector  ports.SSHConnector
	DockerManager docker.DockerManager
}

func NewDockerUseCase(sshAdadpter ports.SSHConnector, DockerManager *docker.DockerManager) *DockerUseCase {
	return &DockerUseCase{
		sshConnector:  sshAdadpter,
		DockerManager: *DockerManager,
	}
}

func (d *DockerUseCase) Execute(envDir string) error {
	var targetIp string
	var err error

	scriptAbsPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, envDir, "deploy.sh")

	if targetIp, err = d.DockerManager.GetTargetIp(scriptAbsPath); err != nil {
		return err
	}
	sshConfig := d.DockerManager.GetSSHConfig(targetIp)
	if err = d.sshConnector.OpenConnection(sshConfig); err != nil {
		log.Fatal("Error opening SSH connection:", err)
	}
	defer d.sshConnector.CloseConnection()

	containers := d.DockerManager.ListContainers()

	selectedContainer := d.DockerManager.SelectContainer(containers)

	d.DockerManager.SelectAndExecuteCommand(selectedContainer)

	return nil
}
