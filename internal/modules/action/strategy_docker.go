package action

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
)

type DockerActionStrategy struct{}

func NewDockerActionStrategy() *DockerActionStrategy {
	return &DockerActionStrategy{}
}

func (s *DockerActionStrategy) GetTemplates() map[string]string {
	const continueMsg = "echo 'Done, press Enter to continue...' && read"
	return map[string]string{
		".host  - bash":                         "bash",
		".host  - htop":                         "htop",
		"docker - logs -f <xyz>":                "docker logs -f %s",
		"docker - exec -u root -it <xyz> sh":    "docker exec -u root -it %s sh",
		"docker - logs <xyz> | less":            "docker logs %s | less",
		"docker - restart <xyz>":                "docker restart %s && " + continueMsg,
		"docker - stats":                        "docker stats",
		"docker - ps    (watch)":                "watch docker ps",
		"docker - ps -a (watch)":                "watch docker ps -a",
		"docker - inspect <xyz> | jq | less":    "docker inspect %s | jq | less",
		"docker - django shell":                 "docker exec -u root -it %s python manage.py shell",
		"docker - django shell_plus":            "docker exec -u root -it %s python manage.py shell_plus",
		"docker - show traefik routes config":   "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq '.routers' | less",
		"docker - show traefik services config": "docker exec %s sh -c 'apk add --no-cache --no-progress -q curl && curl localhost:8080/api/rawdata' | jq '.services' | less",
	}
}

func (s *DockerActionStrategy) GetResources(ssh ports.SSHConnector) ([]string, error) {
	output, err := ssh.RunCommand("docker ps --format json")
	if err != nil {
		return nil, fmt.Errorf("failed to list docker containers: %w", err)
	}

	var resourceNames []string
	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")

	for line := range lines {
		if line == "" {
			continue
		}
		var container struct {
			Names string `json:"Names"`
		}
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			continue
		}
		resourceNames = append(resourceNames, container.Names)
	}

	return resourceNames, nil
}
