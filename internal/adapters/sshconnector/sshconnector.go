package sshconnector

import (
	"fmt"
	"net"
	"time"

	"github.com/goodylabs/tug/internal/ports"
	"golang.org/x/crypto/ssh"
)

type sshConnector struct {
	client *ssh.Client
}

func NewSSHConnector() ports.SSHConnector {
	return &sshConnector{}
}

func (a *sshConnector) OpenConnection(user, host string, port int) error {
	authMethods, err := loadSSHKeysFromDir()
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            user,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return fmt.Errorf("failed to dial SSH: %w", err)
	}

	a.client = client
	return nil
}

func (a *sshConnector) CloseConnection() error {
	if a.client == nil {
		return nil
	}
	err := a.client.Close()
	a.client = nil
	return err
}

func (a *sshConnector) RunCommand(cmd string) (string, error) {
	if a.client == nil {
		return "", fmt.Errorf("connection not opened")
	}
	session, err := a.client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	return string(output), err
}
