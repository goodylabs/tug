package sshconnector

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func loadSSHKeysFromDir() ([]ssh.AuthMethod, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("cannot determine home directory: %w", err)
	}

	sshDir := filepath.Join(homeDir, ".ssh")

	files, err := ioutil.ReadDir(sshDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read ~/.ssh directory: %w", err)
	}

	var authMethods []ssh.AuthMethod

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		path := filepath.Join(sshDir, file.Name())
		keyData, err := os.ReadFile(path)

		if err != nil {
			continue
		}

		signer, err := ssh.ParsePrivateKey(keyData)
		if err != nil {
			continue
		}

		authMethods = append(authMethods, ssh.PublicKeys(signer))
	}

	if len(authMethods) == 0 {
		return nil, fmt.Errorf("no valid SSH private keys found in ~/.ssh/")
	}

	return authMethods, nil
}
