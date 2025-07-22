package adapters

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
	"github.com/goodylabs/docker-swarm-cli/internal/utils"
)

type DockerAdapter struct{}

func (d *DockerAdapter) ConfigureDocker(targetIp string) {
	var dockerHost string
	if !strings.HasPrefix(targetIp, "unix") {
		dockerHost = "ssh://root@" + targetIp
	} else {
		dockerHost = targetIp
	}
	config.DOCKER_HOST_ENV = "DOCKER_HOST=" + dockerHost
}

func (d *DockerAdapter) makeCliCall(args []string) []string {
	cmdArgs := append(args, "--format", "'{{json .}}'")
	cmd := exec.Command("docker", cmdArgs...)

	env := os.Environ()
	env = append(env, config.DOCKER_HOST_ENV)
	cmd.Env = env

	output, err := cmd.Output()
	if err != nil {
		cmdStr := fmt.Sprintf("docker %s", strings.Join(cmdArgs, " "))
		log.Fatal(fmt.Errorf("Error: %s - for command: %s", err, cmdStr))
	}
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	return lines
}

func (d *DockerAdapter) ListServices() []dto.ServiceDTO {
	output := d.makeCliCall([]string{"service", "ls"})
	return utils.UnmarshalLines[dto.ServiceDTO](output)
}

func (d *DockerAdapter) ListContainers() []dto.ContainerDTO {
	output := d.makeCliCall([]string{"ps"})
	return utils.UnmarshalLines[dto.ContainerDTO](output)
}

func NewDockerAdapter() ports.DockerPort {
	return &DockerAdapter{}
}
