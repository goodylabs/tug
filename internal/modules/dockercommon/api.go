package dockercommon

import (
	"errors"

	"github.com/goodylabs/tug/internal/modules/swarm/services"
	"github.com/goodylabs/tug/internal/ports"
)

type EnvCfg struct {
	Name  string
	User  string
	Hosts []string
}

type DockerCommon struct {
	SSHConnector ports.SSHConnector
	Config       map[string]EnvCfg
}

func NewDockerCommon(sshConnector ports.SSHConnector) *DockerCommon {
	return &DockerCommon{
		SSHConnector: sshConnector,
		Config:       make(map[string]EnvCfg),
	}
}

func (d *DockerCommon) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
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

func (d *DockerCommon) GetAvailableActionTemplates() map[string]string {
	return services.GetActionTemplates()
}
