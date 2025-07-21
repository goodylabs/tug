package mocks

import (
	"fmt"

	"github.com/goodylabs/tug/internal/ports"
)

type shellExecutorMock struct {
}

func NewShellExecutorMock(commandsToRun []string) ports.ShellExecutor {
	return &shellExecutorMock{}
}

func (p *shellExecutorMock) Exec(command string) error {
	fmt.Println("Running: ", command)
	return nil
}
