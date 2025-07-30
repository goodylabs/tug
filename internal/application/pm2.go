package application

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/services/pm2"
	"github.com/goodylabs/tug/internal/utils/contextstack"
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

var chunks = [3]string{}

func (p *Pm2UseCase) Execute(envArg string) error {

	var pm2Config dto.EconsystemConfig
	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	if err := p.pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config); err != nil {
		return fmt.Errorf("error loading PM2 config: %w", err)
	}

	contextStack := contextstack.NewContextStack()
	var running = true
	for running == true {
		if contextStack.GetSSHConfig() == nil {
			availableEnvs, err := p.pm2Manager.GetAvailableEnvs(&pm2Config, envArg)
			if err != nil {
				return fmt.Errorf("error while selecting environment: %w", err)
			}

			selectedEnv, err := p.prompter.ChooseFromList(availableEnvs, "Select PM2 <environment>")
			if err != nil {
				running = false
				continue
			}

			fmt.Printf("Connecting via SSH to '%s' server...\n ", selectedEnv)

			sshConfig, err := p.pm2Manager.GetSSHConfig(&pm2Config, selectedEnv)
			if err != nil {
				return fmt.Errorf("getting SSH config: %w", err)
			}

			if err := p.sshConnector.CloseConnection(); err != nil {
				return fmt.Errorf("Error while closing SSH connection: '%w'", err)
			}

			if err := p.sshConnector.OpenConnection(sshConfig); err != nil {
				return fmt.Errorf("Error while opening SSH connection: '%w'", err)
			}

			contextStack.SetSSHConfig(sshConfig)
			continue
		}
		if contextStack.GetResource() == "" {
			resources, err := p.pm2Manager.GetPm2Processes()
			if err != nil {
				return fmt.Errorf("selecting PM2 resource: %w", err)
			}

			resource, err := p.prompter.ChooseFromList(resources, "Select PM2 <resource>")
			if err != nil {
				contextStack.ClearSSHConfig()
				continue
			}
			contextStack.SetResource(resource)
			continue
		}
		if contextStack.GetAction() == "" {
			cmdTemplate, err := p.prompter.ChooseFromMap(pm2.CommandTemplates, "Chose command")
			if err != nil {
				contextStack.ClearResource()
				continue
			}
			contextStack.SetAction(cmdTemplate)
		}
		remoteCmd := fmt.Sprintf(contextStack.GetAction(), contextStack.GetResource())
		if err := p.sshConnector.RunInteractiveCommand(remoteCmd); err != nil {
			contextStack.ClearAction()
			continue
		}
	}
	return nil
}
