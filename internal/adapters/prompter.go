package adapters

import (
	"fmt"
	"sync/atomic"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/choose"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
)

type prompter struct {
	ctrlCPressed atomic.Bool
}

func NewPrompter() ports.Prompter {
	return &prompter{}
}

func (p *prompter) ChooseFromList(options []string, label string) (string, error) {
	if len(options) == 1 {
		return options[0], nil
	}

	utils.SortOptions(options)

	p.clear()
	result, err := prompt.New().
		Ask(label).
		Choose(
			options,
			choose.WithDefaultIndex(0),
			choose.WithHelp(false),
		)
	return result, err
}

func (p *prompter) clear() {
	fmt.Print("\033[H\033[2J")
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
	utils.SortOptions(keys)

	p.clear()
	resultKey, err := prompt.New().
		Ask(label).
		Choose(
			keys,
			choose.WithDefaultIndex(0),
			choose.WithHelp(false),
		)
	if err != nil {
		return "", err
	}

	return options[resultKey], nil
}
