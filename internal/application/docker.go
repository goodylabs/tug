package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/docker"
)

type DockerUseCase struct {
	sshConnector  ports.SSHConnector
	dockerManager *docker.DockerManager
}

func NewDockerUseCase(sshConnector ports.SSHConnector, dockerManager *docker.DockerManager) *DockerUseCase {
	return &DockerUseCase{
		sshConnector:  sshConnector,
		dockerManager: dockerManager,
	}
}

func (d *DockerUseCase) Execute(envDir string) error {
	scriptPath := filepath.Join(config.BASE_DIR, constants.DEVOPS_DIR, envDir, "deploy.sh")

	targetIP, err := d.dockerManager.GetTargetIP(scriptPath)
	if err != nil {
		return fmt.Errorf("getting target IP: %w", err)
	}

	sshConfig := d.dockerManager.GetSSHConfig(targetIP)

	if err := d.sshConnector.ConfigureSSHConnection(sshConfig); err != nil {
		return fmt.Errorf("opening SSH connection: %w", err)
	}
	defer d.sshConnector.CloseConnection()

	containers, err := d.dockerManager.ListContainers()
	if err != nil {
		return fmt.Errorf("listing containers: %w", err)
	}

	container, err := d.dockerManager.SelectContainer(containers)
	if err != nil {
		return fmt.Errorf("selecting container: %w", err)
	}

	if err := d.dockerManager.RunCommandOnContainer(container); err != nil {
		return fmt.Errorf("running command on container: %w", err)
	}

	return nil
}
