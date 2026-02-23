package loadproject

type StrategyName string

const (
	DockerStrategy StrategyName = "docker"
)

type EnvConfig struct {
	Name  string
	User  string
	Hosts []string
}

type ProjectConfig struct {
	Config map[string]EnvConfig
}

type LoadStrategy interface {
	Execute() (ProjectConfig, error)
}
