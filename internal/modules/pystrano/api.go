package pystrano

import (
	"fmt"
	"path/filepath"

	"github.com/goodylabs/tug/internal/modules/pystrano/services"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/config"
)

type envCfg map[string][]string

type PystranoManager struct {
	sshconnector ports.SSHConnector
	config       envCfg
}

func NewPystranoManager(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &PystranoManager{
		sshconnector: sshConnector,
		config:       make(envCfg),
	}
}

func (p *PystranoManager) LoadConfigFromFile() error {
	pystranoDir := filepath.Join(config.GetBaseDir(), "deploy")

	deploymentFiles, err := services.FindDeploymentYAML(pystranoDir)
	if err != nil {
		return err
	}
	if len(deploymentFiles) == 0 {
		return fmt.Errorf("no pystrano config files found in 'deploy' directory")
	}

	for _, file := range deploymentFiles {
		deployFilePath := filepath.Join(pystranoDir, file)
		hosts, err := services.RetrieveHostsFromConfigFile(deployFilePath)
		if err != nil {
			continue
		}
		if len(hosts) == 0 {
			continue
		}
		p.config[file] = hosts
	}

	if len(p.config) == 0 {
		return fmt.Errorf("could not load any valid pystrano config files from 'deploy' directory")
	}

	return nil
}

func (p *PystranoManager) GetAvailableEnvs() ([]string, error) {
	var envs []string
	for env, _ := range p.config {
		envs = append(envs, env)
	}
	return envs, nil
}

func (p *PystranoManager) GetAvailableHosts(env string) ([]string, error) {
	hosts, exists := p.config[env]
	if !exists {
		return []string{}, fmt.Errorf("no such environment: %s", env)
	}
	return hosts, nil
}

func (p *PystranoManager) GetSSHConfig(env, host string) (*ports.SSHConfig, error) {
	hosts, exists := p.config[env]
	if !exists {
		return nil, fmt.Errorf("no such environment: %s", env)
	}

	for _, h := range hosts {
		if h == host {
			return &ports.SSHConfig{
				Host: h,
				Port: 22,
				User: "root",
			}, nil
		}
	}
	return nil, fmt.Errorf("no such host: %s in environment: %s", host, env)
}

func (p *PystranoManager) GetAvailableResources(sshConfig *ports.SSHConfig) ([]string, error) {
	return []string{}, fmt.Errorf("func GetAvailableResources not implemented")
}

func (p *PystranoManager) GetAvailableActionTemplates() map[string]string {
	return services.GetActionTemplates()
}

// func (d *PystranoManager) GetAvailableResources(*ports.SSHConfig) ([]string, error) {
// 	if d.Config == nil {
// 		return []string{}, errors.New("Can not get available resources - config is not loaded")
// 	}

// 	pystranoListCmd := "pystrano ps --format json"
// 	output, err := d.SSHConnector.RunCommand(pystranoListCmd)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return services.GetResourcesFromJsonOutput(output)
// }

// func (d *PystranoManager) GetAvailableActionTemplates() map[string]string {
// 	return services.GetActionTemplates()
// }
