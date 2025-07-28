package adapters

import (
	"fmt"
	"log"
	"sort"

	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/manifoldco/promptui"
)

type prompter struct{}

func NewPrompter() ports.Prompter {
	return &prompter{}
}

func (p *prompter) ChooseFromList(options []string, label string) string {
	fmt.Print("\033[H\033[2J")

	p.sortOptions(options)

	promptUI := promptui.Select{
		Label: label,
		Items: options,
		Size:  10,
	}

	_, result, err := promptUI.Run()
	if err != nil {
		log.Fatal("Error:", err)
	}

	return result
}

func (p *prompter) sortOptions(options []string) {
	sort.Slice(options, func(i, j int) bool {
		if options[i] == constants.ALL_OPTION {
			return false //
		}
		if options[j] == constants.ALL_OPTION {
			return true
		}
		return options[i] < options[j]
	})
}
