package modules

import "github.com/goodylabs/tug/internal/ports"

type TechnologyHandler interface {
	LoadConfigFromFile() error
	GetAvailableEnvs() ([]string, error)
	GetAvailableHosts(env string) ([]string, error)
	GetSSHConfig(env, host string) (*ports.SSHConfig, error)
	GetAvailableResources(sshConfig *ports.SSHConfig) ([]string, error)
	GetAvailableActionTemplates() map[string]string
}

type TechCmdTemplate struct {
	display  string
	template string
	resource string
}
