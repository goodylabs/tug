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
	sshconnector ports.SSHConnector
	prompter     ports.Prompter
}

func NewDockerManager(prompter ports.Prompter, sshconnector ports.SSHConnector) ports.DockerManager {
	return &DockerManager{
		prompter:     prompter,
		sshconnector: sshconnector,
	}
}

func (d *DockerManager) GetTargetIp(scriptAbsPath string) (string, error) {
	fields := []string{
		constants.TARGET_IP_FIELD_LEGACY,
		constants.TARGET_IP_FIELD,
	}

	for _, field := range fields {
		lines, err := utils.GetFileLines(scriptAbsPath)
		if err != nil {
			return "", fmt.Errorf("placeholder Error reading file: %s", scriptAbsPath)
		}

		targetIp := d.getVariableValueFromLines(lines, field)
		if targetIp != "" {
			return targetIp, nil
		}
	}
	return "", errors.New("placeholder")
}

func (d *DockerManager) getVariableValueFromLines(lines []string, key string) string {
	prefix := key + "="
	for _, line := range lines {
		if after, ok := strings.CutPrefix(line, prefix); ok {
			value := after
			value = strings.Trim(value, `"`)
			return value
		}
	}
	return ""
}

func (d *DockerManager) ListContainers() []dto.ContainerDTO {
	var containers []dto.ContainerDTO
	output, err := d.sshconnector.RunCommand("docker ps --format json")
	if err != nil {
		fmt.Println("Error running docker ps command:", err)
	}

	lines := strings.SplitSeq(strings.TrimSpace(output), "\n")
	for line := range lines {
		var container dto.ContainerDTO
		if err := json.Unmarshal([]byte(line), &container); err != nil {
			fmt.Println("Error unmarshalling line:", err)
			continue
		}
		containers = append(containers, container)
	}

	return containers
}

func (d *DockerManager) ChoseContainer(containers []dto.ContainerDTO) dto.ContainerDTO {
	var names []string
	for _, container := range containers {
		names = append(names, container.Name)
	}
	selectedName := d.prompter.ChooseFromList(names, "Chose container")
	for _, container := range containers {
		if selectedName == container.Name {
			return container
		}
	}
	panic(constants.PANIC)
}

func (d *DockerManager) SelectAndExecuteCommand(container dto.ContainerDTO) {
	commands := []string{
		"docker exec -it " + container.Name + " bash",
		"docker logs " + container.Name,
		"docker stop " + container.Name,
		"docker start " + container.Name,
	}

	selectedCommand := d.prompter.ChooseFromList(commands, "Select command to execute")
	if selectedCommand == "" {
		fmt.Println("No command selected, exiting.")
		return
	}

	d.sshconnector.RunInteractiveCommand(selectedCommand)
}
