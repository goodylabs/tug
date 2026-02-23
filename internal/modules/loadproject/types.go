package loadproject

type ProjectConfig struct {
	Config map[string]EnvConfig
}

type EnvConfig struct {
	Name  string
	User  string
	Hosts []string
}

type GetEnvironments struct {
	Config ProjectConfig
}
