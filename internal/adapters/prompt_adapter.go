package adapters

import (
	"fmt"
	"log"

	"github.com/goodylabs/docker-swarm-cli/internal/config"
	"github.com/goodylabs/docker-swarm-cli/internal/ports"
	"github.com/manifoldco/promptui"
)

type PromptAdapter struct{}

func (p *PromptAdapter) ChooseFromList(options []string, label string) (string, error) {
	fmt.Print("\033[H\033[2J")

	promptUI := promptui.Select{
		Label: label,
		Items: options,
	}

	_, result, err := promptUI.Run()
	if err != nil {
		log.Fatal("Error:", err)
	}

	return result, nil
}

func NewPromptAdapter() ports.PromptPort {
	if config.TESTING {
		return &MockPromptAdapter{}
	}
	return &PromptAdapter{}
}
