package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
)

type PystranoActionStrategy struct{}

func NewPystranoActionStrategy() *PystranoActionStrategy {
	return &PystranoActionStrategy{}
}

func (s *PystranoActionStrategy) GetTemplates() map[string]string {
	const continueMsg = "echo 'Done, press Enter to continue...' && read"

	return map[string]string{
		".host     - bash":          "bash",
		".host     - htop":          "htop",
		"pystrano  - logs -f <xyz>": "pystrano logs -f %s",
		"pystrano  - restart <xyz>": "pystrano restart %s && " + continueMsg,
		"pystrano  - status":        "pystrano ps",
	}
}

func (s *PystranoActionStrategy) GetResources(ssh ports.SSHConnector) ([]string, error) {
	output, err := ssh.RunCommand("pystrano ps --format json")
	if err != nil {
		return nil, fmt.Errorf("failed to list pystrano resources: %w", err)
	}

	type pystranoResource struct {
		Name string `json:"Name"`
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	var resourceNames []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "{") {
			continue
		}
		var res pystranoResource
		if err := json.Unmarshal([]byte(line), &res); err != nil {
			continue
		}
		if res.Name != "" {
			resourceNames = append(resourceNames, res.Name)
		}
	}

	return resourceNames, nil
}
