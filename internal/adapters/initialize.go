package adapters

import (
	"github.com/goodylabs/tug/internal/adapters/dockercli"
	"github.com/goodylabs/tug/internal/adapters/httpconnector"
	"github.com/goodylabs/tug/internal/adapters/prompter"
	"github.com/goodylabs/tug/internal/adapters/shellexecutor"
	"github.com/goodylabs/tug/internal/ports"
)

var DockerApi ports.DockerApi
var Prompter ports.Prompter
var ShellExecutor ports.ShellExecutor
var HttpConnector ports.HttpConnector

func InitializeDependencies(options ...func()) {
	if DockerApi == nil {
		DockerApi = dockercli.NewDockerApi()
	}
	if Prompter == nil {
		Prompter = prompter.NewPrompter()
	}
	if ShellExecutor == nil {
		ShellExecutor = shellexecutor.NewShellExecutor()
	}

	if HttpConnector == nil {
		HttpConnector = httpconnector.NewHttpConnector()
	}

	for _, opt := range options {
		opt()
	}
}

func WithDockerApi(dockerApi ports.DockerApi) func() {
	return func() {
		DockerApi = dockerApi
	}
}

func WithPrompter(prompter ports.Prompter) func() {
	return func() {
		Prompter = prompter
	}
}

func WithShellExecutor(shellExecutor ports.ShellExecutor) func() {
	return func() {
		ShellExecutor = shellExecutor
	}
}
