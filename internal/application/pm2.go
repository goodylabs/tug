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
	pm2Manager   ports.Pm2Manager
	sshconnector ports.SSHConnector
	prompter     ports.Prompter
}

func NewPm2UseCase(pm2Manager ports.Pm2Manager, sshconnector ports.SSHConnector, prompter ports.Prompter) *Pm2UseCase {
	return &Pm2UseCase{
		pm2Manager:   pm2Manager,
		sshconnector: sshconnector,
		prompter:     prompter,
	}
}

func (p *Pm2UseCase) Execute() {
	var pm2ConfigDTO dto.EconsystemConfigDTO

	ecosystemConfigPath := filepath.Join(config.BASE_DIR, constants.ECOSYSTEM_CONFIG_FILE)
	err := p.pm2Manager.LoadPm2Config(ecosystemConfigPath, &pm2ConfigDTO)
	if err != nil {
		log.Fatal("Error loading PM2 config:", err)
	}

	fmt.Println(pm2ConfigDTO.Deploy["staging"].Host)

	// TODO wybór enva
	selectedEnv := "staging"
	// TODO wybór hosta
	hostIndex := 0
	sshUser := pm2ConfigDTO.Deploy[selectedEnv].User
	sshHost := pm2ConfigDTO.Deploy[selectedEnv].Host[hostIndex]

	err = p.sshconnector.OpenConnection(sshUser, sshHost, 22)
	if err != nil {
		log.Fatal("Error opening SSH connection:", err)
	}
	defer p.sshconnector.CloseConnection()

	selectedResource := p.selectResource()
	fmt.Println(selectedResource)
}

func (p *Pm2UseCase) selectResource() string {
	output, err := p.sshconnector.RunCommand(pm2.JLIST_CMD)
	if err != nil {
		log.Fatalf("Error running PM2 command: %v", err)
	}

	var pm2List []dto.Pm2ListItemDTO
	err = p.pm2Manager.JsonOutputHandler(output, &pm2List)
	if err != nil {
		log.Fatalf("Error parsing PM2 list output: %v", err)
	}
	return pm2List[0].Name
}

// zaczytanie configa
// ssh connection -> defer close
//
// wybranie resource (jest)
//
// wybranie komendy (jest)
