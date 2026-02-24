package action

import (
	"fmt"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/utils"
)

type SSHService struct {
	connector ports.SSHConnector
}

func NewSSHService(connector ports.SSHConnector) *SSHService {
	return &SSHService{connector: connector}
}

// GetConnector wystawia interfejs SSH do zadań specjalnych (np. listowanie zasobów)
func (s *SSHService) GetConnector() ports.SSHConnector {
	return s.connector
}

func (s *SSHService) Connect(user, host string) (string, error) {
	cfg := &ports.SSHConfig{
		Host: host,
		User: user,
		Port: 22,
	}

	if err := s.connector.ConfigureSSHConnection(cfg); err != nil {
		errorMsg := fmt.Errorf("failed to connect as %s@%s: %w", cfg.User, cfg.Host, err)
		return "", errorMsg
	}

	hostname, err := s.connector.RunCommand("hostname")
	if err != nil {
		return "unknown_host", nil
	}
	return utils.NormalizeSpaces(hostname), nil
}

func (s *SSHService) RunAction(template, resource string) {
	cmd := template
	if resource != "" {
		cmd = fmt.Sprintf(template, resource)
	}
	s.connector.RunInteractiveCommand(cmd)
}
