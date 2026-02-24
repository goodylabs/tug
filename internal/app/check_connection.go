package app

import (
	"github.com/goodylabs/tug/internal/modules/checkconnections"
	"github.com/goodylabs/tug/internal/modules/loadproject"
)

type CheckConnectionUseCase struct {
	checkConnectionsService *checkconnections.CheckConnectionsService
}

func NewCheckConnectionUseCase() *CheckConnectionUseCase {
	return &CheckConnectionUseCase{
		checkConnectionsService: checkconnections.NewCheckConnectionsService(),
	}
}

func (p *CheckConnectionUseCase) Execute(loadTech loadproject.StrategyName) error {
	lp := loadproject.NewLoadProject()
	pCfg, err := lp.Execute(loadTech)
	if err != nil {
		return err
	}

	p.checkConnectionsService.Execute(pCfg)
	return nil
}
