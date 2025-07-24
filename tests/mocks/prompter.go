package mocks

import (
	"github.com/goodylabs/tug/internal/ports"
)

type prompterMock struct {
	choicesInOrder []int
	currentIndex   int
}

func NewPrompterMock(choicesInOrder []int) ports.Prompter {
	return &prompterMock{
		choicesInOrder: choicesInOrder,
	}
}

func (p *prompterMock) ChooseFromList(options []string, label string) string {
	index := p.choicesInOrder[p.currentIndex]
	p.currentIndex++
	return options[index]
}

func (p *prompterMock) ChooseByName(list any, label string) any {
	return map[string]any{}
}
