package ports

import "github.com/goodylabs/tug/internal/dto"

type Prompter interface {
	ChooseFromList([]string, string) (string, error)
	ChooseFromMap(map[string]string, string) (string, error)
}

type SSHConnector interface {
	ConfigureSSHConnection(*dto.SSHConfig) error
	CloseConnection() error
	RunCommand(cmd string) (string, error)
	RunInteractiveCommand(cmd string) error
}
