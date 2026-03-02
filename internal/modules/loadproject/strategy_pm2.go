package loadproject

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/internal/modules"
	"github.com/goodylabs/tug/pkg/config"
)

const (
	tmpJsonPath        = "/tmp/tug_pm2_config.json"
	ecosystemJsScript  = `const fs = require('fs'); const config = require('%s'); fs.writeFileSync('%s', JSON.stringify(config));`
	ecosystemCjsScript = `const fs = require('fs'); const config = require('%s'); fs.writeFileSync('%s', JSON.stringify(config));`
)

type FlexHost []string

func (fh *FlexHost) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*fh = []string{single}
		return nil
	}
	var slice []string
	if err := json.Unmarshal(data, &slice); err != nil {
		return err
	}
	*fh = slice
	return nil
}

type pm2ConfigDTO struct {
	Deploy map[string]struct {
		User string   `json:"user"`
		Host FlexHost `json:"host"`
	} `json:"deploy"`
}

type Pm2LoadStrategy struct{}

func NewPm2LoadStrategy() *Pm2LoadStrategy {
	return &Pm2LoadStrategy{}
}

func (s *Pm2LoadStrategy) Execute() (modules.ProjectConfig, error) {
	projectCfg := modules.ProjectConfig{
		Config: make(map[string]modules.EnvConfig),
	}

	configPath, err := s.getPm2ConfigPath(config.GetBaseDir())
	if err != nil {
		return projectCfg, err
	}

	if err := s.convertJsFileToJson(configPath); err != nil {
		return projectCfg, err
	}
	defer os.Remove(tmpJsonPath)

	jsonFile, err := os.ReadFile(tmpJsonPath)
	if err != nil {
		return projectCfg, fmt.Errorf("failed to read temp PM2 json: %w", err)
	}

	var dto pm2ConfigDTO
	if err := json.Unmarshal(jsonFile, &dto); err != nil {
		return projectCfg, fmt.Errorf("failed to unmarshal PM2 config: %w", err)
	}

	for envName, deployCfg := range dto.Deploy {
		if deployCfg.User == "" || len(deployCfg.Host) == 0 {
			continue
		}

		projectCfg.Config[envName] = modules.EnvConfig{
			Name:  envName,
			User:  deployCfg.User,
			Hosts: deployCfg.Host,
		}
	}

	return projectCfg, nil
}

func (s *Pm2LoadStrategy) getPm2ConfigPath(dir string) (string, error) {
	options := []string{"ecosystem.config.cjs", "ecosystem.config.js"}
	for _, name := range options {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("ecosystem config file not found in %s", dir)
}

func (s *Pm2LoadStrategy) convertJsFileToJson(path string) error {
	var nodeScript string
	switch {
	case strings.HasSuffix(path, ".js"):
		nodeScript = fmt.Sprintf(ecosystemJsScript, path, tmpJsonPath)
	case strings.HasSuffix(path, ".cjs"):
		nodeScript = fmt.Sprintf(ecosystemCjsScript, path, tmpJsonPath)
	default:
		return fmt.Errorf("unsupported file extension: %s", path)
	}

	cmd := exec.Command("node", "-e", nodeScript)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("node conversion failed: %w (output: %s)", err, string(output))
	}
	return nil
}
