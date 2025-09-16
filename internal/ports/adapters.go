package ports

type PromptOptions struct {
	Value string
	Label string
}

type Prompter interface {
	ChooseFromList([]string, string) (string, error)
	ChooseFromMap(map[string]string, string) (string, error)
	ChooseFromListWithDisplayValue([]PromptOptions, string) (string, error)
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
