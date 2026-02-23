package app

import (
	"fmt"

	"github.com/goodylabs/tug/internal/modules/loadproject"
)

type UseModuleV2UseCase struct{}

func NewUseModuleV2UseCase() *UseModuleV2UseCase {
	return &UseModuleV2UseCase{}
}

func (u *UseModuleV2UseCase) Execute(
	loadProjectStrategy loadproject.StrategyName,
	// actionStrategy
) error {
	lp := loadproject.NewLoadProject()

	projectCfg, err := lp.Execute(loadProjectStrategy)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", projectCfg)

	return nil
}
