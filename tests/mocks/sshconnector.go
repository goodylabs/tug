package mocks

import (
	"bytes"
	"os"
	"os/exec"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
)

type sshconnectorMock struct{}

func NewSSHConnectorMock() ports.SSHConnector {
	return &sshconnectorMock{}
}

func (m *sshconnectorMock) OpenConnection(sshConfig *dto.SSHConfigDTO) error {
	return nil
}

func (m *sshconnectorMock) CloseConnection() error {
	return nil
}

func (m *sshconnectorMock) RunCommand(cmd string) (string, error) {
	command := exec.Command("sh", "-c", cmd)
	var out bytes.Buffer
	command.Stdout = &out
	command.Stderr = &out
	err := command.Run()
	return out.String(), err
}

func (m *sshconnectorMock) RunInteractiveCommand(cmd string) error {
	command := exec.Command("sh", "-c", cmd)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
