package adapters

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/goodylabs/tug/internal/ports"
	"github.com/manifoldco/promptui"
	"golang.org/x/term"
)

type prompter struct {
	lastIndexes map[string]int
}

func NewPrompter() ports.Prompter {
	return &prompter{
		lastIndexes: make(map[string]int),
	}
}

func (p *prompter) ChooseFromList(options []string, label string) (string, error) {
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

func (p *prompter) hashOptions(options []ports.DisplayValueOpts) string {
	labels := make([]string, len(options))
	for i, opt := range options {
		labels[i] = opt.Label
	}
	hash := sha256.Sum256([]byte(strings.Join(labels, "|")))
	return hex.EncodeToString(hash[:])
}

type noBellWriter struct {
	w io.Writer
}

func (n noBellWriter) Write(p []byte) (int, error) {
	filtered := make([]byte, 0, len(p))
	for _, b := range p {
		if b != 0x07 {
			filtered = append(filtered, b)
		}
	}
	return n.w.Write(filtered)
}

func (n noBellWriter) Close() error {
	return nil
}

func (p *prompter) runPrompter(options []ports.DisplayValueOpts, label string) (string, error) {
	p.clear()

	sort.Slice(options, func(i, j int) bool {
		return options[i].Label < options[j].Label
	})

	_, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(err)
	}

	hashKey := p.hashOptions(options)
	lastIndex := p.lastIndexes[hashKey]

	prompt := promptui.Select{
		Label:             label,
		Items:             options,
		Size:              height - 3,
		StartInSearchMode: false,
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
		CursorPos: lastIndex,
		Stdout:    noBellWriter{os.Stdout}, // tu filtrujemy BEL
	}

	i, _, err := prompt.Run()
	p.clear()
	if err != nil {
		return "", err
	}

	p.lastIndexes[hashKey] = i
	return options[i].Value, nil
}

func (p *prompter) ChooseFromListWithDisplayValue(options []ports.DisplayValueOpts, label string) (string, error) {
	p.clear()
	return p.runPrompter(options, label)
}
