package adapters

import "errors"

type MockPromptAdapter struct{}

func (p *MockPromptAdapter) ChooseFromList(options []string, label string) (string, error) {
	for _, option := range options {
		if option == "localhost" {
			return option, nil
		}
	}
	return "", errors.New("You f*cked something up with the test environment setup")
}
