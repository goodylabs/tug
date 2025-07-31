package adapters

import (
	"fmt"
	"os"
	"strings"

	// "github.com/cqroot/prompt"
	// "github.com/cqroot/prompt/choose"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

type prompter struct{}

func NewPrompter() ports.Prompter {
	return &prompter{}
}

func (p *prompter) ChooseFromList(options []string, label string) (string, error) {
	if len(options) == 1 {
		return options[0], nil
	}

	p.clear()

	return p.runPrompter(options, label)
}

func (p *prompter) ChooseFromMap(options map[string]string, label string) (string, error) {
	if len(options) == 1 {
		for _, v := range options {
			return v, nil
		}
	}

	keys := make([]string, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}

	p.clear()

	resultKey, err := p.runPrompter(keys, label)

	if err != nil {
		return "", err
	}

	return options[resultKey], nil
}

func (p *prompter) clear() {
	fmt.Print("\033[H\033[2J")
}

func (p *prompter) runPrompter(options []string, label string) (string, error) {
	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	utils.SortOptions(options)

	prompt := promptui.Select{
		Label:             label,
		Items:             options,
		Size:              height - 3,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return strings.Contains(options[index], input)
		},
	}

	_, result, err := prompt.Run()
	p.clear()
	if err != nil {
		return "", err
	}
	return result, nil
}
