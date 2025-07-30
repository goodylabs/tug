package mocks

import (
	"errors"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
)

type sshConnectorMock struct {
	runCmdOutput           string
	runCmdErr              error
	expectedInteractiveCmd string
}

func NewSSHConnectorMock(runCmdOutput string, runCmdErr error) ports.SSHConnector {
	return &sshConnectorMock{
		runCmdOutput: runCmdOutput,
		runCmdErr:    runCmdErr,
	}
}

func NewSSHConnectorInteractiveCommandMock(expectedInteractiveCmd string) ports.SSHConnector {
	return &sshConnectorMock{
		expectedInteractiveCmd: expectedInteractiveCmd,
	}
}

func (m *sshConnectorMock) ConfigureSSHConnection(sshConfig *dto.SSHConfig) error {
	return nil
}

func (m *sshConnectorMock) CloseConnection() error {
	return nil
}

func (m *sshConnectorMock) RunCommand(cmd string) (string, error) {
	return m.runCmdOutput, m.runCmdErr
}

func (m *sshConnectorMock) RunInteractiveCommand(cmd string) error {
	if cmd != m.expectedInteractiveCmd {
		return errors.New("unexpected interactive command: " + cmd)
	}
	return nil
}
