package loadproject

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goodylabs/tug/pkg/config"
)

const (
	tmpJsonPath = "/tmp/tug_pm2_config.json"
	// Skrypty wymagają dwóch parametrów: %s (wejście) i %s (wyjście)
	ecosystemJsScript  = `const fs = require('fs'); const config = require('%s'); fs.writeFileSync('%s', JSON.stringify(config));`
	ecosystemCjsScript = `const fs = require('fs'); const config = require('%s'); fs.writeFileSync('%s', JSON.stringify(config));`
)

type Pm2LoadStrategy struct{}

func NewPm2LoadStrategy() *Pm2LoadStrategy {
	return &Pm2LoadStrategy{}
}

func (s *Pm2LoadStrategy) Execute() (ProjectConfig, error) {
	projectCfg := ProjectConfig{
		Config: make(map[string]EnvConfig),
	}

	configPath, err := s.getPm2ConfigPath(config.GetBaseDir())
	if err != nil {
		return projectCfg, err
	}

	if err := s.convertJsFileToJson(configPath); err != nil {
		return projectCfg, err
	}

	// Odczyt JSONa wygenerowanego przez Node.js
	jsonFile, err := os.ReadFile(tmpJsonPath)
	if err != nil {
		return projectCfg, fmt.Errorf("failed to read temp PM2 json: %w", err)
	}
	defer os.Remove(tmpJsonPath) // Czyścimy /tmp

	var dto pm2ConfigDTO
	if err := json.Unmarshal(jsonFile, &dto); err != nil {
		return projectCfg, fmt.Errorf("failed to unmarshal PM2 config: %w", err)
	}

	// Mapujemy dane z PM2 na nasz wspólny format ProjectConfig
	for envName, deployCfg := range dto.Deploy {
		if deployCfg.User == "" || len(deployCfg.Host) == 0 {
			continue
		}

		projectCfg.Config[envName] = EnvConfig{
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

// Struktura pomocnicza do Unmarshal
type pm2ConfigDTO struct {
	Deploy map[string]struct {
		User string   `json:"user"`
		Host []string `json:"host"`
	} `json:"deploy"`
}
