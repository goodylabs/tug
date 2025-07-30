package adapters

import (
	"errors"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/goodylabs/tug/internal/dto"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/tughelper"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type sshConnector struct {
	client *ssh.Client
}

func NewSSHConnector() ports.SSHConnector {
	return &sshConnector{}
}

func (s *sshConnector) ConfigureSSHConnection(sshConfig *dto.SSHConfig) error {
	authMethods, err := s.loadSSHKeysFromDir()
	if err != nil {
		return err
	}

	config := &ssh.ClientConfig{
		User:            sshConfig.User,
		Auth:            authMethods,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	address := net.JoinHostPort(sshConfig.Host, fmt.Sprintf("%d", sshConfig.Port))
	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return err
	}

	s.client = client
	return nil
}

func (s *sshConnector) CloseConnection() error {
	if s.client == nil {
		return nil
	}
	err := s.client.Close()
	s.client = nil
	return err
}

func (s *sshConnector) RunCommand(cmd string) (string, error) {
	if s.client == nil {
		return "", nil
	}
	session, err := s.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	return string(output), err
}

func (s *sshConnector) RunInteractiveCommand(cmd string) error {
	if s.client == nil {
		return nil
	}

	session, err := s.client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		return err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState)

	return session.Run(cmd)
}

func (s *sshConnector) loadSSHKeysFromDir() ([]ssh.AuthMethod, error) {
	tugConfig, err := tughelper.GetTugConfig()
	if err != nil {
		return []ssh.AuthMethod{}, errors.New("Can not read tug config file, run `tug initialize` to configure tug.")
	}

	keyData, err := os.ReadFile(tugConfig.SSHKeyPath)
	if err != nil {
		return []ssh.AuthMethod{}, err
	}
	signer, err := ssh.ParsePrivateKey(keyData)
	if err != nil {
		return []ssh.AuthMethod{}, err
	}
	return []ssh.AuthMethod{ssh.PublicKeys(signer)}, nil
}
