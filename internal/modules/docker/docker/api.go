package docker

import (
	"errors"

	"github.com/goodylabs/tug/internal/modules"
	dockercommon "github.com/goodylabs/tug/internal/modules/docker/common"
	"github.com/goodylabs/tug/internal/modules/docker/docker/services"
	"github.com/goodylabs/tug/internal/ports"
)

type DockerManager struct {
	*dockercommon.DockerCommon
}

func NewDockerManager(sshConnector ports.SSHConnector) modules.TechnologyHandler {
	return &DockerManager{
		DockerCommon: dockercommon.NewDockerCommon(sshConnector),
	}
}

func (d *DockerManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	if d.DockerCommon.Config == nil {
		return []string{}, errors.New("Can not get available resources - config is not loaded")
	}

	dockerListCmd := "docker ps --format json"
	output, err := d.DockerCommon.SSHConnector.RunCommand(dockerListCmd)
	if err != nil {
		return nil, err
	}

	return services.GetResourcesFromJsonOutput(output)
}

func (d *DockerManager) GetAvailableActionTemplates() map[string]string {
	return services.GetActionTemplates()
}
