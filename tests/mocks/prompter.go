package mocks

import (
	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
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
