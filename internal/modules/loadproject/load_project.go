package loadproject

import "fmt"

type LoadProject struct{}

func NewLoadProject() *LoadProject {
	return &LoadProject{}
}

func (lp *LoadProject) Execute(name StrategyName) (ProjectConfig, error) {
	var strategy LoadStrategy

	switch name {
	case DockerStrategy:
		strategy = NewDockerLoadStrategy()
	// Tutaj w przyszłości dodasz:
	// case K8sStrategy:
	//     strategy = NewK8sLoadStrategy()
	default:
		return ProjectConfig{}, fmt.Errorf("unsupported strategy: %s", name)
	}

	return strategy.Execute()
}
