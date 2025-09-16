package modules

import (
	"github.com/goodylabs/tug/internal/ports"
)

type TechnologyHandler interface {
	LoadConfigFromFile() error
	GetAvailableEnvs() ([]string, error)
	GetAvailableHosts(env string) ([]string, error)
	GetSSHConfig(env, host string) (*ports.SSHConfig, error)
	GetAvailableResources(sshConfig *ports.SSHConfig) ([]string, error)
	GetAvailableActionTemplates() []TechCmdTemplate
}

type TechCmdTemplate struct {
	Display  string
	Template string
	Resource string
}
