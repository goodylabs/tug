package mocks

import (
	"fmt"

	"github.com/goodylabs/tug/internal/ports"
)

type shellExecutor struct{}

func NewShellExecutor() ports.ShellExecutor {
	return &shellExecutor{}
}

func (p *shellExecutor) Exec(command string) error {
	fmt.Println("Runnig command:", command)
	return nil
}
