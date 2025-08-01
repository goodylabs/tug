package pm2

type pm2ConfigDTO struct {
	Apps []struct {
		Name string `json:"name"`
	} `json:"apps"`
	Deploy map[string]struct {
		User string   `json:"user"`
		Host []string `json:"host"`
	} `json:"deploy"`
}

func (e *pm2ConfigDTO) ListEnvironments() []string {
	envs := make([]string, 0, len(e.Deploy))
	for env := range e.Deploy {
		envs = append(envs, env)
	}
	return envs
}

func (e *pm2ConfigDTO) ListHostsInEnv(env string) []string {
	hosts := []string{}
	for host := range e.Deploy[env].Host {
		hosts = append(hosts, e.Deploy[env].Host[host])
	}
	return hosts
}

type pm2ListItemDTO struct {
	Name string `json:"name"`
}
