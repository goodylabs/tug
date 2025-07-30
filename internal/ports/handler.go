package ports

import "github.com/goodylabs/tug/internal/dto"

type TechnologyHandler interface {
	GetAvailableEnvs() ([]string, error)
	GetSSHConfig(env string) (*dto.SSHConfig, error)
	// GetAvailableActions() ([]string, error)
	GetAvailableResources(*dto.SSHConfig) ([]string, error)
	// ExecuteAction(action string, resource string, env string) error
}
