package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
)

type SwarmActionStrategy struct{}

func NewSwarmActionStrategy() *SwarmActionStrategy {
	return &SwarmActionStrategy{}
}

func (s *SwarmActionStrategy) GetTemplates() map[string]string {
	const continueMsg = "echo 'Done, press Enter to continue...' && read"

	return map[string]string{
		".host  - bash":                             "bash",
		".host  - htop":                             "htop",
		"swarm - ls (watch)":                        "watch docker service ls",
		"swarm - ps <xyz> (watch)":                  "watch docker service ps %s --no-trunc",
		"swarm - ps <xyz> [ only running ] (watch)": `watch 'docker service ps --filter desired-state=running --format "{{.ID}} {{.Name}} - {{.Node}} | {{.Image}}" %s'`,
		"swarm - inspect <xyz> | jq | less":         "docker service inspect %s | jq | less",
		"swarm - restart <xyz>":                     "docker service update %s --force && " + continueMsg,
		"swarm - logs -f <xyz>":                     "docker service logs -f %s",
		"swarm - logs <xyz> | less":                 "docker service logs %s | less",
		"swarm - scale <xyz> to 0":                  "docker service scale %s=0 && " + continueMsg,
		"swarm - scale <xyz> to 1":                  "docker service scale %s=1 && " + continueMsg,
		"swarm - scale <xyz> to 3":                  "docker service scale %s=3 && " + continueMsg,
	}
}

func (s *SwarmActionStrategy) GetResources(ssh ports.SSHConnector) ([]string, error) {

	output, err := ssh.RunCommand("docker service ls --format json")
	if err != nil {
		return nil, fmt.Errorf("failed to list docker services: %w", err)
	}

	type serviceDTO struct {
		Name string `json:"Name"`
	}

	var resourceNames []string

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		var service serviceDTO
		if err := json.Unmarshal([]byte(line), &service); err != nil {

			continue
		}

		if service.Name != "" {
			resourceNames = append(resourceNames, service.Name)
		}
	}

	return resourceNames, nil
}
