package ports

type PromptPort interface {
	ChooseFromList(options []string, label string) (string, error)
}
