package docker

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

type DockerManager struct {
	sshConnector ports.SSHConnector
	prompter     ports.Prompter
}

func NewDockerManager(prompter ports.Prompter, sshConnector ports.SSHConnector) *DockerManager {
	return &DockerManager{
		prompter:     prompter,
		sshConnector: sshConnector,
	}
}

func (d *DockerManager) GetTargetIP(scriptPath string) (string, error) {
	for _, field := range []string{constants.TARGET_IP_FIELD_LEGACY, constants.TARGET_IP_FIELD} {
		lines, err := utils.GetFileLines(scriptPath)
		if err != nil {
			return "", fmt.Errorf("reading deploy.sh: %w", err)
		}

		if ip := extractVariable(lines, field); ip != "" {
			return ip, nil
		}
	}
	return "", errors.New("could not find TARGET_IP in deploy.sh")
}

func extractVariable(lines []string, key string) string {
	prefix := key + "="
	for _, line := range lines {
		if value, ok := strings.CutPrefix(line, prefix); ok {
			return strings.Trim(value, `"`)
		}
	}
	return ""
}

func (d *DockerManager) GetSSHConfig(ip string) *dto.SSHConfigDTO {
	return &dto.SSHConfigDTO{
		User: "root",
		Host: ip,
		Port: 22,
	}
}

func (d *DockerManager) ListContainers() ([]dto.ContainerDTO, error) {
	var containers []dto.ContainerDTO

	output, err := d.sshConnector.RunCommand("docker ps --format json")
	if err != nil {
		return nil, fmt.Errorf("docker ps failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(output), "\n")
	for _, line := range lines {
		var container dto.ContainerDTO
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			continue
		}
		containers = append(containers, container)
	}

	return containers, nil
}

func (d *DockerManager) SelectContainer(containers []dto.ContainerDTO) (dto.ContainerDTO, error) {
	if len(containers) == 0 {
		return dto.ContainerDTO{}, errors.New("no containers available")
	}

	var names []string
	for _, c := range containers {
		names = append(names, c.Name)
	}

	selected := d.prompter.ChooseFromList(names, "Choose container")
	for _, c := range containers {
		if c.Name == selected {
			return c, nil
		}
	}
	return dto.ContainerDTO{}, fmt.Errorf("container %s not found", selected)
}

func (d *DockerManager) RunCommandOnContainer(container dto.ContainerDTO) error {
	commands := []string{
		"docker logs -f " + container.Name,
		"docker exec -it " + container.Name + " sh",
		"docker logs " + container.Name,
		"docker stop " + container.Name,
		"docker start " + container.Name,
	}

	selected := d.prompter.ChooseFromList(commands, "Select command to execute")
	if selected == "" {
		return errors.New("no command selected")
	}

	if err := d.sshConnector.RunInteractiveCommand(selected); err != nil {
		return fmt.Errorf("executing command failed: %w", err)
	}

	return nil
}
