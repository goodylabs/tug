package action

import "fmt"

type StrategyName string

const (
	Docker   StrategyName = "docker"
	Pm2      StrategyName = "pm2"
	Pystrano StrategyName = "pystrano"
	Swarm    StrategyName = "swarm"
)

func GetStrategy(tech StrategyName) (ActionStrategy, error) {
	switch tech {
	case Docker:
		return NewDockerActionStrategy(), nil
	case Pm2:
		return NewPm2ActionStrategy(), nil
	// case Pystrano:
	// 	return NewPystranoActionStrategy(), nil
	// case Swarm:
	//     return NewSwarmActionStrategy(), nil
	default:
		return nil, fmt.Errorf("unsupported action strategy: %s", tech)
	}
}
