package pm2

import (
	"encoding/json"
	"errors"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
)

type Pm2Handler struct {
	pm2Config    *pm2ConfigDTO
	sshConnector ports.SSHConnector
}

func NewPm2Handler(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &Pm2Handler{
		sshConnector: sshConnector,
	}
}

func (p *Pm2Handler) LoadConfigFromFile() error {
	if p.pm2Config != nil {
		return errors.New("Can not load config - it is already loaded")
	}

	configPath, err := getPm2ConfigPath()
	if err != nil {
		return err
	}

	if err := convertJsFileToJson(configPath); err != nil {
		return err
	}

	return nil
}

func (p *Pm2Handler) GetAvailableEnvs() ([]string, error) {
	if p.pm2Config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	if len(p.pm2Config.Deploy) == 0 {
		return []string{}, errors.New("No environments found in PM2 config")
	}

	var options []string
	for env := range p.pm2Config.Deploy {
		options = append(options, env)
	}
	return options, nil
}

func (p *Pm2Handler) GetAvailableHosts(env string) ([]string, error) {
	if p.pm2Config == nil {
		return nil, errors.New("Can not get available hosts - config is not loaded")
	}

	hosts := p.pm2Config.ListHostsInEnv(env)
	if len(hosts) == 0 {
		return nil, errors.New("No hosts found for the specified environment")
	}

	return hosts, nil
}

func (p *Pm2Handler) GetSSHConfig(env, host string) (*dto.SSHConfig, error) {
	if p.pm2Config == nil {
		return nil, errors.New("Can not get SSH config - PM2 config is not loaded")
	}

	envConfig, exists := p.pm2Config.Deploy[env]
	if !exists {
		return nil, errors.New("Environment not found in PM2 config")
	}

	return &dto.SSHConfig{
		User: envConfig.User,
		Host: host,
		Port: 22,
	}, nil
}

func (p *Pm2Handler) GetAvailableResources(sshConfig *dto.SSHConfig) ([]string, error) {
	output, err := p.sshConnector.RunCommand(jlistCmd)
	if err != nil {
		return nil, err
	}

	var pm2List []pm2ListItemDTO
	if err := json.Unmarshal([]byte(output), &pm2List); err != nil {
		return nil, errors.New("failed to parse PM2 list output: " + err.Error())
	}

	var resources []string
	for _, item := range pm2List {
		resources = append(resources, item.Name)
	}

	return resources, nil
}

func (p *Pm2Handler) GetAvailableActionTemplates() map[string]string {
	return commandTemplates
}
