package swarm

type dockerConfigEnv struct {
	Name  string
	User  string
	Hosts []string
}

type dockerConfig struct {
	Envs map[string]dockerConfigEnv
}

type serviceDTO struct {
	ID       string `json:"ID"`
	Image    string `json:"Image"`
	Mode     string `json:"Mode"`
	Name     string `json:"Name"`
	Ports    string `json:"Ports"`
	Replicas string `json:"Replicas"`
}
