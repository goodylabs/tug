package adapters

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

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

func (s *sshConnector) ConfigureSSHConnection(sshConfig *ports.SSHConfig) error {
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

func (s *sshConnector) loadSSHKeysFromDir() ([]ssh.AuthMethod, error) {
	tugConfig, err := tughelper.GetTugConfig()
	if err != nil {
		return nil, errors.New("Can not read tug config file, run `tug initialize` to configure tug.")
	}

	keyData, err := os.ReadFile(tugConfig.SSHKeyPath)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(keyData)
	if err != nil {
		if _, ok := err.(*ssh.PassphraseMissingError); ok {
			fmt.Print("Enter passphrase for SSH key: ")
			passphrase, perr := term.ReadPassword(int(os.Stdin.Fd()))
			fmt.Println()
			if perr != nil {
				return nil, perr
			}

			signer, err = ssh.ParsePrivateKeyWithPassphrase(keyData, passphrase)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	return []ssh.AuthMethod{ssh.PublicKeys(signer)}, nil
}

func (s *sshConnector) RunCommand(cmd string) (string, error) {
	var ErrSSHConnection = fmt.Errorf("Error establishing SSH connection to the server.")
	if s.client == nil {
		return "", ErrSSHConnection
	}
	session, err := s.client.NewSession()
	if err != nil {
		return "", ErrSSHConnection
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	return string(output), nil
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

	fd := int(os.Stdin.Fd())
	width, height, err := term.GetSize(fd)
	if err != nil {
		return err
	}

	if err := session.RequestPty("xterm", height, width, modes); err != nil {
		return err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer term.Restore(fd, oldState)

	sigWinch := make(chan os.Signal, 1)
	signal.Notify(sigWinch, syscall.SIGWINCH)
	defer signal.Stop(sigWinch)

	go func() {
		for range sigWinch {
			width, height, err := term.GetSize(fd)
			if err != nil {
				continue
			}
			session.WindowChange(height, width)
		}
	}()

	return session.Run(cmd)
}
