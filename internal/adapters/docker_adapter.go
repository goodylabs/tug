package adapters

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/goodylabs/docker-swarm-cli/internal/dto"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
)

type DockerAdapter struct {
	dockerHost string
}

func (d *DockerAdapter) ConfigureDocker(targetIp string) {
	if !strings.HasPrefix(targetIp, "unix") {
		d.dockerHost = "ssh://root@" + targetIp
	} else {
		d.dockerHost = targetIp
	}
}

func (d *DockerAdapter) makeCliCall(args []string) []string {
	cmdArgs := append(args, "--format", "'{{json .}}'")
	cmd := exec.Command("docker", cmdArgs...)

	if d.dockerHost != "" {
		env := os.Environ()
		env = append(env, "DOCKER_HOST="+d.dockerHost)
		cmd.Env = env
	}

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
	return unmarshalLines[dto.ServiceDTO](output)
}

func (d *DockerAdapter) ListContainers() []dto.ContainerDTO {
	output := d.makeCliCall([]string{"ps"})
	return unmarshalLines[dto.ContainerDTO](output)
}

func unmarshalLines[T any](output []string) []T {
	var result []T
	for _, line := range output {
		var item T
		line = strings.Trim(line, "'")
		err := json.Unmarshal([]byte(line), &item)
		if err != nil {
			log.Printf("failed to unmarshal line: %s, err: %v", line, err)
			continue
		}
		result = append(result, item)
	}
	return result
}

func NewDockerAdapter() ports.DockerPort {
	return &DockerAdapter{}
}
