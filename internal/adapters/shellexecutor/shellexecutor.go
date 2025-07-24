package shellexecutor

import (
	"log"
	"os"
	"os/exec"

	"github.com/goodylabs/tug/internal/ports"
)

type shellExecutor struct{}

func NewShellExecutor() ports.ShellExecutor {
	return &shellExecutor{}
}

func (s *shellExecutor) Exec(command string) error {
	shell := os.Getenv("SHELL")
	if shell == "" {
		log.Fatal("placeholder")
	}
	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
