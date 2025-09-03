package services

import (
	"encoding/json"
	"strings"
)

type containerDTO struct {
	Command      string `json:"Command"`
	CreatedAt    string `json:"CreatedAt"`
	ID           string `json:"ID"`
	Image        string `json:"Image"`
	Labels       string `json:"Labels"`
	LocalVolumes string `json:"LocalVolumes"`
	Mounts       string `json:"Mounts"`
	Name         string `json:"Names"`
	Networks     string `json:"Networks"`
	Ports        string `json:"Ports"`
	RunningFor   string `json:"RunningFor"`
	Size         string `json:"Size"`
	State        string `json:"State"`
	Status       string `json:"Status"`
}

func GetResourcesFromJsonOutput(output string) ([]string, error) {

	var containers []containerDTO

	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")
	for line := range lines {
		var container containerDTO
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			continue
		}
		containers = append(containers, container)
	}

	var containerNames []string
	for _, container := range containers {
		containerNames = append(containerNames, container.Name)
	}
	return containerNames, nil
}
