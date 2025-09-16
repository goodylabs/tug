package swarm

import (
	"errors"

	"github.com/goodylabs/tug/internal/modules"
	dockercommon "github.com/goodylabs/tug/internal/modules/docker/common"
	"github.com/goodylabs/tug/internal/modules/docker/swarm/services"
	"github.com/goodylabs/tug/internal/ports"
)

type SwarmManager struct {
	*dockercommon.DockerCommon
}

func NewSwarmManager(sshConnector ports.SSHConnector) modules.TechnologyHandler {
	return &SwarmManager{
		DockerCommon: dockercommon.NewDockerCommon(sshConnector),
	}
}

func (d *SwarmManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
	if d.Config == nil {
		return []string{}, errors.New("Can not get available resources - config is not loaded")
	}

	dockerListCmd := "docker service ls --format json"
	output, err := d.SSHConnector.RunCommand(dockerListCmd)
	if err != nil {
		return nil, err
	}

	return services.GetResourcesFromJsonOutput(output)
}

func (d *SwarmManager) GetAvailableActionTemplates() []modules.TechCmdTemplate {
	return services.GetActionTemplates()
}
