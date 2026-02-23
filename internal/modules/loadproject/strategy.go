package loadproject

type LoadStrategy interface {
	Execute() (ProjectConfig, error)
}

type ProjectConfig struct {
	Config map[string]EnvConfig
}

type EnvConfig struct {
	Name  string
	User  string
	Hosts []string
}
