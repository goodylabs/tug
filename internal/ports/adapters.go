package ports

type Prompter interface {
	ChooseFromList([]string, string) string
}

type SSHConnector interface {
	OpenConnection(user, host string, port int) error
	CloseConnection() error
	RunCommand(cmd string) (string, error)
	RunInteractiveCommand(cmd string) error
}
