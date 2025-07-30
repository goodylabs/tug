package ports

import "github.com/goodylabs/tug/internal/dto"

type Prompter interface {
	ChooseFromList([]string, string) (string, error)
	// AskForInput(string) string
}

type SSHConnector interface {
	OpenConnection(*dto.SSHConfig) error
	CloseConnection() error
	RunCommand(cmd string) (string, error)
	RunInteractiveCommand(cmd string) error
}
