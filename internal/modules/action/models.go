package action

import (
	"github.com/goodylabs/tug/internal/ports"
)

type ActionStrategy interface {
	GetResources(ssh ports.SSHConnector) ([]string, error)
	GetTemplates() map[string]string
}

type ActionManager struct {
	strategy ActionStrategy
}

func NewActionManager(strategy ActionStrategy) *ActionManager {
	return &ActionManager{strategy: strategy}
}

func (m *ActionManager) GetAvailableResources(ssh ports.SSHConnector) ([]string, error) {
	return m.strategy.GetResources(ssh)
}

func (m *ActionManager) GetAvailableActionTemplates() map[string]string {
	return m.strategy.GetTemplates()
}
