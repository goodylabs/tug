package dto

type EconsystemConfig struct {
	Apps []struct {
		Name string `json:"name"`
	} `json:"apps"`
	Deploy map[string]struct {
		User string   `json:"user"`
		Host []string `json:"host"`
	} `json:"deploy"`
}

func (e *EconsystemConfig) ListEnvironments() []string {
	envs := make([]string, 0, len(e.Deploy))
	for env := range e.Deploy {
		envs = append(envs, env)
	}
	return envs
}

type Pm2ListItem struct {
	Name string `json:"name"`
}
