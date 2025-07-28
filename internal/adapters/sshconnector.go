package adapters

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"golang.org/x/crypto/ssh"
)

type sshConnector struct {
	client *ssh.Client
}

func NewSSHConnector() ports.SSHConnector {
	return &sshConnector{}
}

func (a *sshConnector) OpenConnection(sshConfig *dto.SSHConfigDTO) error {
	authMethods, err := loadSSHKeysFromDir()
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            sshConfig.User,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	address := net.JoinHostPort(sshConfig.Host, fmt.Sprintf("%d", sshConfig.Port))
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

func (s *sshConnector) RunInteractiveCommand(cmd string) error {
	if s.client == nil {
		return fmt.Errorf("connection not opened")
	}

	session, err := s.client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return fmt.Errorf("request for pseudo terminal failed: %w", err)
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return session.Run(cmd)
}

func loadSSHKeysFromDir() ([]ssh.AuthMethod, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home directory: %w", err)
	}

	possibleKeys := []string{
		"id_rsa",
		"id_ed25519",
		"id_ecdsa",
		"id_dsa",
	}

	for _, keyFile := range possibleKeys {
		path := filepath.Join(homeDir, ".ssh", keyFile)
		keyData, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		signer, err := ssh.ParsePrivateKey(keyData)
		if err != nil {
			continue
		}
		return []ssh.AuthMethod{ssh.PublicKeys(signer)}, nil
	}

	return nil, fmt.Errorf("no valid default SSH private key found in ~/.ssh/")
}
