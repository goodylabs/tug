package application

import (
	"fmt"
	"log"
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

func (p *Pm2UseCase) Execute() error {
	var pm2Config dto.EconsystemConfigDTO

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	if err := p.pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2Config); err != nil {
		return fmt.Errorf("Error loading PM2 config: %v", err)
	}

	selectedEnv := p.pm2Manager.SelectEnvironment(&pm2Config)

	sshConfig := p.pm2Manager.GetSSHConfig(&pm2Config, selectedEnv)

	if err := p.sshConnector.OpenConnection(sshConfig); err != nil {
		log.Fatal("Error opening SSH connection:", err)
	}
	defer p.sshConnector.CloseConnection()

	selectedResource := p.pm2Manager.SelectResource()

	fmt.Println(selectedResource)

	return nil
}

// zaczytanie configa
// ssh connection -> defer close
//
// wybranie resource (jest)
//
// wybranie komendy (jest)
