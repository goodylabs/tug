package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
)

const (
	cmdLogs        string = "Logs (live)"
	cmdShell       string = "Shell"
	cmdDjangoShell string = "Django shell"
)

var AllCommandTypes = []string{
	string(cmdLogs),
	string(cmdShell),
	string(cmdDjangoShell),
}

func SelectCommandsService(prompter ports.PromptPort, container dto.ContainerDTO) {
	choice := prompter.ChooseFromList(AllCommandTypes, "What do you want?")

	command := []string{}

	command = append(command, config.DOCKER_HOST_ENV)

	switch choice {
	case cmdLogs:
		command = append(command, "docker", "logs", container.ID, "-f")
	case cmdShell:
		command = append(command, "docker", "exec", "-u", "root", "-ti", container.ID, "sh")
	case cmdDjangoShell:
		command = append(command, "docker", "exec", "-u", "root", "-ti", container.ID, "python", "manage.py", "shell_plus")
	}

	fullCmd := strings.Join(command, " ")
	runCommand(fullCmd)
}

func runCommand(fullCmd string) {
	// fmt.Println(fullCmd)

	cmd := exec.Command("bash", "-c", fullCmd)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
