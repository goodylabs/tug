package application

import (
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

	contextStack := contextstack.NewGeneric(p.pm2Manager, p.sshConnector, p.prompter)
	return contextStack.Execute()
}
