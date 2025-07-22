package adapters

import (
	"log"

	"github.com/goodylabs/docker-swarm-cli/internal/constants"
)

type MockPromptAdapter struct{}

func (p *MockPromptAdapter) ChooseFromList(options []string, label string) string {
	for _, option := range options {
		if option == "localhost" {
			return option
		}
	}
	log.Fatal("You f*cked something up with the test environment setup")
	panic(constants.PANIC)
}

func (p *MockPromptAdapter) ChooseByName(list any, label string) any {
	return map[string]any{}
}
