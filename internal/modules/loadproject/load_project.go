package loadproject

import (
	"fmt"

	"github.com/goodylabs/tug/internal/modules"
)

type StrategyName string

const (
	DockerStrategy   StrategyName = "docker"
	Pm2Strategy      StrategyName = "pm2"
	PystranoStrategy StrategyName = "pystrano"
	InputStrategy    StrategyName = "input"
)

type LoadProject struct{}

func NewLoadProject() *LoadProject {
	return &LoadProject{}
}

func (lp *LoadProject) Execute(strategyName StrategyName) (modules.ProjectConfig, error) {
	var strategy modules.LoadStrategy

	switch strategyName {
	case DockerStrategy:
		strategy = NewDockerLoadStrategy()
	case Pm2Strategy:
		strategy = NewPm2LoadStrategy()
	case PystranoStrategy:
		strategy = NewPystranoLoadStrategy()
	case InputStrategy:
		strategy = NewInputLoadStrategy()
	default:
		return modules.ProjectConfig{}, fmt.Errorf("unsupported strategy: %s", strategyName)
	}

	return strategy.Execute()
}
