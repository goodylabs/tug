package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/pm2"
)

type Pm2UseCase struct {
	pm2Manager   *pm2.Pm2Manager
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
}

func NewPm2UseCase(pm2Manager *pm2.Pm2Manager, sshConnector ports.SSHConnector, prompter ports.Prompter) *Pm2UseCase {
	return &Pm2UseCase{
		pm2Manager:   pm2Manager,
		sshConnector: sshConnector,
		prompter:     prompter,
	}
}

func (p *Pm2UseCase) Execute(envArg string) error {
	var pm2Config dto.EconsystemConfig

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	if err := p.pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config); err != nil {
		return fmt.Errorf("error loading PM2 config: %w", err)
	}

	selectedEnv, err := p.pm2Manager.SelectEnvFromConfig(&pm2Config, envArg)
	if err != nil {
		return fmt.Errorf("error while selecting environment: %w", err)
	}

	fmt.Printf("Connecting via SSH to '%s' server...\n ", selectedEnv)

	sshConfig, err := p.pm2Manager.GetSSHConfig(&pm2Config, selectedEnv)
	if err != nil {
		return fmt.Errorf("getting SSH config: %w", err)
	}

	if err := p.sshConnector.OpenConnection(sshConfig); err != nil {
		return fmt.Errorf("Error while opening SSH connection: '%w'", err)
	}
	defer p.sshConnector.CloseConnection()

	resource, err := p.pm2Manager.SelectResource()
	if err != nil {
		return fmt.Errorf("selecting PM2 resource: %w", err)
	}

	cmdTemplate, err := p.pm2Manager.SelectCommandTemplate()
	if err != nil {
		return fmt.Errorf("Error while selecting command template: %w", err)
	}

	remoteCmd := fmt.Sprintf(cmdTemplate, resource)

	return p.sshConnector.RunInteractiveCommand(remoteCmd)
}
