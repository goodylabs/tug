package loadproject

import (
	"fmt"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/modules"
	"github.com/goodylabs/tug/internal/ports"
)

type InputLoadStrategy struct {
	prompter ports.Prompter
}

func NewInputLoadStrategy() *InputLoadStrategy {
	return &InputLoadStrategy{
		prompter: adapters.NewPrompter(),
	}
}

func (s *InputLoadStrategy) Execute() (modules.ProjectConfig, error) {
	host, err := s.prompter.AskUserForInput("Enter IP address")
	if err != nil {
		return modules.ProjectConfig{}, err
	}

	defaultUser := "root"
	userPrompt := fmt.Sprintf("Enter username (default: %s)", defaultUser)
	user, err := s.prompter.AskUserForInput(userPrompt)
	if err != nil {
		return modules.ProjectConfig{}, err
	}

	if user == "" {
		user = defaultUser
	}

	defaultEnv := map[string]modules.EnvConfig{
		"default": {
			User: user,
			Hosts: []string{
				host,
			}},
	}

	return modules.ProjectConfig{
		Config: defaultEnv,
	}, nil
}
