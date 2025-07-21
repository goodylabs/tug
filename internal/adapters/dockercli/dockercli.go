package dockercli

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

type dockerCli struct{}

func NewDockerApi() ports.DockerApi {
	return &dockerCli{}
}

func (d *dockerCli) ConfigureDocker(targetIp string) {
	if !strings.HasPrefix(targetIp, "unix") {
		os.Setenv("DOCKER_HOST", "ssh://root@"+targetIp)
	}
}

func (d *dockerCli) makeCliCallOutputJson(args []string) []string {
	allArgs := append(args, "--format", "{{ json . }}")
	cmd := exec.Command("docker", allArgs...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		cmdStr := fmt.Sprintf("docker %s", strings.Join(args, " "))
		wrappedErr := fmt.Errorf("failed: %s\nerror: %w\noutput:\n%s", cmdStr, err, string(output))
		log.Fatal(wrappedErr)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	return lines
}

func (d *dockerCli) ListServices() []dto.ServiceDTO {
	output := d.makeCliCallOutputJson([]string{"service", "ls"})
	return utils.UnmarshalLines[dto.ServiceDTO](output)
}

func (d *dockerCli) ListContainers() []dto.ContainerDTO {
	output := d.makeCliCallOutputJson([]string{"ps"})
	return utils.UnmarshalLines[dto.ContainerDTO](output)
}

func (d *dockerCli) execCmdWithStdio(args []string) error {
	cmd := exec.Command("docker", args...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (d *dockerCli) execCmd(args []string) error {
	cmd := exec.Command("docker", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func (d *dockerCli) ExecLogsLive(containerId string) {
	err := d.execCmdWithStdio([]string{"logs", containerId, "-f"})
	if err != nil {
		log.Fatal("placeholder")
	}
}

func (d *dockerCli) ExecShell(containerId string) {
	err := d.execCmdWithStdio([]string{"exec", "-u", "root", "-ti", containerId, "sh"})
	if err != nil {
		log.Fatal("placeholder")
	}
}

func (d *dockerCli) ExecDjangoShell(containerId string) {
	err := d.execCmdWithStdio([]string{"docker", "exec", "-u", "root", "-ti", containerId, "python", "manage.py", "shell_plus"})
	if err != nil {
		log.Fatal("placeholder")
	}
}

func (d *dockerCli) ExecLogs(containerId string) {
	err := d.execCmd([]string{"logs", containerId})
	if err != nil {
		log.Fatal("placeholder")
	}
}
