package swarm

import (
	"errors"

	"github.com/goodylabs/tug/internal/modules/dockercommon"
	"github.com/goodylabs/tug/internal/modules/swarm/services"
	"github.com/goodylabs/tug/internal/ports"
)

type SwarmManager struct {
	*dockercommon.DockerCommon
}

func NewSwarmManager(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &SwarmManager{
		DockerCommon: dockercommon.NewDockerCommon(sshConnector),
	}
}

func (d *SwarmManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	if d.Config == nil {
		return []string{}, errors.New("Can not get available resources - config is not loaded")
	}

	dockerListCmd := "docker ps --format json"
	output, err := d.SSHConnector.RunCommand(dockerListCmd)
	if err != nil {
		return nil, err
	}

	return services.GetResourcesFromJsonOutput(output)
}

func (d *SwarmManager) GetAvailableActionTemplates() map[string]string {
	return services.GetActionTemplates()
}
