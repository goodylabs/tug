package ports

import "github.com/goodylabs/tug/internal/dto"

type TechnologyHandler interface {
	GetAvailableEnvs() ([]string, error)
	GetSSHConfig(env string) (*dto.SSHConfig, error)
	GetAvailableResources(*dto.SSHConfig) ([]string, error)
}
