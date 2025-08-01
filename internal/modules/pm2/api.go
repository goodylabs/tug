package pm2

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
)

type Pm2Handler struct {
	config       *pm2ConfigDTO
	sshConnector ports.SSHConnector
}

func NewPm2Handler(sshConnector ports.SSHConnector) ports.TechnologyHandler {
	return &Pm2Handler{
		sshConnector: sshConnector,
	}
}

func (p *Pm2Handler) LoadConfigFromFile() error {
	if p.config != nil {
		return errors.New("Can not load config - it is already loaded")
	}

	configPath, err := GetPm2ConfigPath(config.BASE_DIR)
	if err != nil {
		return err
	}

	if err := ConvertJsFileToJson(configPath); err != nil {
		return err
	}

	jsonFile, err := os.ReadFile(tmpJsonPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonFile, &p.config); err != nil {
		return errors.New("failed to unmarshal PM2 config: " + err.Error())
	}

	if len(p.config.Deploy) == 0 {
		return errors.New("PM2 config is missing 'deploy' environments")
	}

	for env, cfg := range p.config.Deploy {
		if cfg.User == "" {
			return errors.New("missing user for environment: " + env)
		}
		if len(cfg.Host) == 0 {
			return errors.New("no hosts defined for environment: " + env)
		}
	}

	return nil
}

func (p *Pm2Handler) GetAvailableEnvs() ([]string, error) {
	if p.config == nil {
		return []string{}, errors.New("Can not get available environments - config is not loaded")
	}

	return p.config.ListEnvironments(), nil
}

func (p *Pm2Handler) GetAvailableHosts(env string) ([]string, error) {
	if p.config == nil {
		return nil, errors.New("Can not get available hosts - config is not loaded")
	}

	return p.config.ListHostsInEnv(env), nil
}

func (p *Pm2Handler) GetSSHConfig(env, host string) (*dto.SSHConfig, error) {
	if p.config == nil {
		return nil, errors.New("Can not get SSH config - PM2 config is not loaded")
	}

	return &dto.SSHConfig{
		User: p.config.Deploy[env].User,
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
