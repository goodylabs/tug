package mocks

import (
	"os"
	"os/exec"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
)

type sshConnectorMock struct {
	runCmdOutput string
	runCmdErr    error
}

func NewSSHConnectorMock(runCmdOutput string, runCmdErr error) ports.SSHConnector {
	return &sshConnectorMock{
		runCmdOutput: runCmdOutput,
		runCmdErr:    runCmdErr,
	}
}

func (m *sshConnectorMock) OpenConnection(sshConfig *dto.SSHConfigDTO) error {
	return nil
}

func (m *sshConnectorMock) CloseConnection() error {
	return nil
}

func (m *sshConnectorMock) RunCommand(cmd string) (string, error) {
	return m.runCmdOutput, m.runCmdErr
}

func (m *sshConnectorMock) RunInteractiveCommand(cmd string) error {
	command := exec.Command("sh", "-c", cmd)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
