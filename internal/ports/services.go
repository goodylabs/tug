package ports

import "github.com/goodylabs/tug/internal/dto"

type TechnologyHandler interface {
	LoadConfigFromFile() error
	GetAvailableEnvs() ([]string, error)
	GetAvailableHosts(env string) ([]string, error)
	GetSSHConfig(env, host string) (*dto.SSHConfig, error)
	GetAvailableResources(sshConfig *dto.SSHConfig) ([]string, error)
	GetAvailableActionTemplates() map[string]string
}
