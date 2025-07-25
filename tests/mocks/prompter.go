package mocks

import "github.com/goodylabs/tug/internal/ports"

type prompterMock struct {
	seq   []int
	index int
}

func NewPrompterMock(seq []int) ports.Prompter {
	return &prompterMock{seq: seq}
}

func (p *prompterMock) ChooseFromList(options []string, label string) string {
	i := p.seq[p.index]
	p.index++
	return options[i]
}
