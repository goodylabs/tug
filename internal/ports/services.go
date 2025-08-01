package ports

import "github.com/goodylabs/tug/internal/dto"

type TechnologyHandler interface {
	LoadConfigFromFile() error
	GetAvailableEnvs() ([]string, error)
	GetAvailableHosts(env string) ([]string, error)
	GetSSHConfig(string, string) (*dto.SSHConfig, error)
	GetAvailableResources(*dto.SSHConfig) ([]string, error)
	GetAvailableActionTemplates() map[string]string
}
