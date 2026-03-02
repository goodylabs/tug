package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
)

type Pm2ActionStrategy struct{}

func NewPm2ActionStrategy() *Pm2ActionStrategy {
	return &Pm2ActionStrategy{}
}

func (s *Pm2ActionStrategy) GetTemplates() map[string]string {
	const nvmPrefix = "source ~/.nvm/nvm.sh; "
	const continueMsg = " && echo 'Done, press Enter to continue...' && read"

	return map[string]string{
		".host - bash (nvm)":        nvmPrefix + "bash",
		".host - htop (nvm)":        nvmPrefix + "htop",
		"pm2 - logs <xyz>":          nvmPrefix + "pm2 logs %s",
		"pm2 - logs <xyz> | less":   nvmPrefix + "pm2 logs %s | less",
		"pm2 - logs (all)":          nvmPrefix + "pm2 logs",
		"pm2 - status <xyz> | less": nvmPrefix + "pm2 status %s | less",
		"pm2 - show <xyz>":          nvmPrefix + "pm2 show %s" + continueMsg,
		"pm2 - describe <xyz>":      nvmPrefix + "pm2 describe %s" + continueMsg,
		"pm2 - restart <xyz>":       nvmPrefix + "pm2 restart %s",
		"pm2 - monit":               nvmPrefix + "pm2 monit",
		"pm2 - update":              nvmPrefix + "pm2 update",
	}
}

func (s *Pm2ActionStrategy) GetResources(ssh ports.SSHConnector) ([]string, error) {
	const jlistCmd = `source ~/.nvm/nvm.sh; pm2 jlist | sed -n '/^\[/,$p'`

	output, err := ssh.RunCommand(jlistCmd)
	if err != nil {
		return nil, fmt.Errorf("failed to list pm2 processes: %w", err)
	}

	type pm2Item struct {
		Name string `json:"name"`
	}

	var pm2List []pm2Item
	if err := json.Unmarshal([]byte(strings.TrimSpace(output)), &pm2List); err != nil {
		return nil, fmt.Errorf("failed to parse PM2 list output: %w", err)
	}

	var resourceNames []string
	for _, item := range pm2List {
		if item.Name != "" {
			resourceNames = append(resourceNames, item.Name)
		}
	}

	return resourceNames, nil
}
