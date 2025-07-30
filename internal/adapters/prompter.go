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
	fmt.Println("\033[H\033[2J")

	utils.SortOptions(options)

	result, err := prompt.New().
		Ask(label).
		Choose(
			options,
			choose.WithDefaultIndex(1),
			choose.WithHelp(true),
		)
	return result, err
}
