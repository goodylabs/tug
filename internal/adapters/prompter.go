package adapters

import (
	"fmt"
	"os"
	"sort"
	"strings"

	// "github.com/cqroot/prompt"
	// "github.com/cqroot/prompt/choose"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

type prompter struct{}

func NewPrompter() ports.Prompter {
	return &prompter{}
}

func (p *prompter) ChooseFromList(options []string, label string) (string, error) {
	p.clear()

	optionsDisplayValueOpts := make([]ports.DisplayValueOpts, len(options))
	for i, key := range options {
		optionsDisplayValueOpts[i] = ports.DisplayValueOpts{
			Label: key,
			Value: key,
		}
	}

	return p.runPrompter(optionsDisplayValueOpts, label)
}

func (p *prompter) ChooseFromMap(options map[string]string, label string) (string, error) {
	keys := make([]string, 0, len(options))
	for k := range options {
		keys = append(keys, k)
	}

	p.clear()

	optionsDisplayValueOpts := make([]ports.DisplayValueOpts, len(keys))
	for i, key := range keys {
		optionsDisplayValueOpts[i] = ports.DisplayValueOpts{
			Label: key,
			Value: key,
		}
	}
	resultKey, err := p.runPrompter(optionsDisplayValueOpts, label)

	if err != nil {
		return "", err
	}

	return options[resultKey], nil
}

func (p *prompter) clear() {
	fmt.Print("\033[H\033[2J")
}

func (p *prompter) runPrompter(options []ports.DisplayValueOpts, label string) (string, error) {
	sort.Slice(options, func(i, j int) bool {
		return options[i].Label < options[j].Label
	})

	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	prompt := promptui.Select{
		Label:             label,
		Items:             options,
		Size:              height - 3,
		StartInSearchMode: true,
		Templates: &promptui.SelectTemplates{
			Label:    "{{ . }}",
			Active:   "▸ {{ .Label | cyan }}",
			Inactive: "  {{ .Label }}",
			Selected: "✔ {{ .Label | green }}",
		},
		Searcher: func(input string, index int) bool {
			option := options[index]
			return strings.Contains(option.Label, input)
		},
	}

	i, _, err := prompt.Run()
	if err != nil {
		p.clear()
		return "", err
	}
	p.clear()
	return options[i].Value, nil
}

func (p *prompter) ChooseFromListWithDisplayValue(options []ports.DisplayValueOpts, label string) (string, error) {
	p.clear()

	return p.runPrompter(options, label)
}
