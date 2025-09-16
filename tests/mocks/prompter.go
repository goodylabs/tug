package mocks

import (
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/pkg/utils"
)

type prompterMock struct {
	seq   []int
	index int
}

func NewPrompterMock(seq []int) ports.Prompter {
	return &prompterMock{
		seq: seq,
	}
}

func (p *prompterMock) ChooseFromList(options []string, label string) (string, error) {
	i := p.seq[p.index]
	p.index++
	utils.SortOptions(options)
	return options[i], nil
}

func (p *prompterMock) ChooseFromMap(map[string]string, string) (string, error) {
	return "", nil
}

func (p *prompterMock) ChooseFromListWithDisplayValue([]ports.PromptOptions, string) (string, error) {
	return "", nil
}
