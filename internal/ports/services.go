package ports

type SSHConfig struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type TechnologyHandler interface {
	LoadConfigFromFile() error
	GetAvailableEnvs() ([]string, error)
	GetAvailableHosts(env string) ([]string, error)
	GetSSHConfig(env, host string) (*SSHConfig, error)
	GetAvailableResources(sshConfig *SSHConfig) ([]string, error)
	GetAvailableActionTemplates() map[string]string
}
