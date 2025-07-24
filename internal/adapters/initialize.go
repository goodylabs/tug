package adapters

import (
	"github.com/goodylabs/tug/internal/adapters/dockercli"
	"github.com/goodylabs/tug/internal/adapters/prompter"
	"github.com/goodylabs/tug/internal/adapters/shellexecutor"
	"github.com/goodylabs/tug/internal/ports"
)

var DockerApi ports.DockerApi
var Prompter ports.Prompter
var ShellExecutor ports.ShellExecutor

func InitializeDependencies(options ...func()) {
	DockerApi = dockercli.NewDockerApi()
	Prompter = prompter.NewPrompter()
	ShellExecutor = shellexecutor.NewShellExecutor()

	for _, opt := range options {
		opt()
	}
}

func WithDockerApi(api ports.DockerApi) func() {
	return func() {
		DockerApi = api
	}
}

func WithPrompter(p ports.Prompter) func() {
	return func() {
		Prompter = p
	}
}

func WithShellExecutor(p ports.ShellExecutor) func() {
	return func() {
		ShellExecutor = p
	}
}
