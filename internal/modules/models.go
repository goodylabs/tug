package modules

import (
	"fmt"
	"slices"
	"sort"
)

type LoadStrategy interface {
	Execute() (ProjectConfig, error)
}

type EnvConfig struct {
	Name  string
	User  string
	Hosts []string
}

type ProjectConfig struct {
	Config map[string]EnvConfig
}

func (pc *ProjectConfig) GetAvailableEnvs() []string {
	envs := make([]string, 0, len(pc.Config))
	for env := range pc.Config {
		envs = append(envs, env)
	}
	sort.Strings(envs)
	return envs
}

func (pc *ProjectConfig) GetAvailableHosts(env string) ([]string, error) {
	cfg, ok := pc.Config[env]
	if !ok {
		return nil, fmt.Errorf("environment not found: %s", env)
	}
	return cfg.Hosts, nil
}

func (pc *ProjectConfig) GetEnvConfig(env string) (EnvConfig, error) {
	cfg, ok := pc.Config[env]
	if !ok {
		return EnvConfig{}, fmt.Errorf("environment not found: %s", env)
	}
	return cfg, nil
}

func (pc *ProjectConfig) IsHostInEnv(env, host string) bool {
	cfg, ok := pc.Config[env]
	if !ok {
		return false
	}
	return slices.Contains(cfg.Hosts, host)
}
