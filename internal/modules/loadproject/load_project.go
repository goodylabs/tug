package loadproject

import "fmt"

type StrategyName string

const (
	DockerStrategy   StrategyName = "docker"
	Pm2Strategy      StrategyName = "pm2"
	PystranoStrategy StrategyName = "pystrano"
)

type LoadProject struct{}

func NewLoadProject() *LoadProject {
	return &LoadProject{}
}

func (lp *LoadProject) Execute(strategyName StrategyName) (ProjectConfig, error) {
	var strategy LoadStrategy

	switch strategyName {
	case DockerStrategy:
		strategy = NewDockerLoadStrategy()
	case Pm2Strategy:
		strategy = NewPm2LoadStrategy()
	case PystranoStrategy:
		strategy = NewPystranoLoadStrategy()
	default:
		return ProjectConfig{}, fmt.Errorf("unsupported strategy: %s", strategyName)
	}

	return strategy.Execute()
}
