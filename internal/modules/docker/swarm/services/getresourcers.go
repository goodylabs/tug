package services

import (
	"encoding/json"
	"strings"
)

type serviceDTO struct {
	ID       string `json:"ID"`
	Image    string `json:"Image"`
	Mode     string `json:"Mode"`
	Name     string `json:"Name"`
	Ports    string `json:"Ports"`
	Replicas string `json:"Replicas"`
}

func GetResourcesFromJsonOutput(output string) ([]string, error) {

	var services []serviceDTO

	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")
	for line := range lines {
		var service serviceDTO
		if err := json.Unmarshal([]byte(line), &service); err != nil {
			continue
		}
		services = append(services, service)
	}

	var serviceNames []string
	for _, service := range services {
		serviceNames = append(serviceNames, service.Name)
	}
	return serviceNames, nil
}
