package ports

type PromptOption struct {
	Value string
	Label string
}

type Prompter interface {
	ChooseFromList([]string, string) (string, error)
	ChooseFromListWithDisplayValue([]PromptOption, string) (string, error)
}

type SSHConfig struct {
	User string `json:"user"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

type SSHConnector interface {
	ConfigureSSHConnection(*SSHConfig) error
	CloseConnection() error
	RunCommand(cmd string) (string, error)
	RunInteractiveCommand(cmd string) error
}
