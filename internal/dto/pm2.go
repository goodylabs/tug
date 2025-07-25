package dto

type EconsystemConfigDTO struct {
	Apps []struct {
		Name string `json:"name"`
	} `json:"apps"`
	Deploy map[string]struct {
		User string   `json:"user"`
		Host []string `json:"host"`
	} `json:"deploy"`
}

func (e *EconsystemConfigDTO) ListEnvironments() []string {
	envs := make([]string, 0, len(e.Deploy))
	for env := range e.Deploy {
		envs = append(envs, env)
	}
	return envs
}
