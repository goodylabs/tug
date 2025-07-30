package adapters

import (
	"fmt"
	"log"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/goodylabs/tug/internal/utils"
	"github.com/manifoldco/promptui"
)

type prompter struct{}

func NewPrompter() ports.Prompter {
	return &prompter{}
}

func (p *prompter) ChooseFromList(options []string, label string) string {
	fmt.Print("\033[H\033[2J")

	utils.SortOptions(options)

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

// func AskForInput(label string) string {
// 	fmt.Print("\033[H\033[2J")

// 	promptUI := promptui.Prompt{
// 		Label: label,
// 	}

// 	result, err := promptUI.Run()
// 	if err != nil {
// 		log.Fatal("Error:", err)
// 	}

// 	return result
// }
