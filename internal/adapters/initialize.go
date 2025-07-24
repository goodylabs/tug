package adapters

import (
	"fmt"

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
	fmt.Println("Initializing dependencies")

	if DockerApi == nil {
		DockerApi = dockercli.NewDockerApi()
		fmt.Println("Initialized prod:DockerApi")
	}
	if Prompter == nil {
		Prompter = prompter.NewPrompter()
		fmt.Println("Initialized prod:Prompter")
	}
	if ShellExecutor == nil {
		ShellExecutor = shellexecutor.NewShellExecutor()
		fmt.Println("Initialized prod:ShellExecutor")
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
		fmt.Println("Initialized test:DockerApi")
	}
}

func WithPrompter(prompter ports.Prompter) func() {
	return func() {
		Prompter = prompter
		fmt.Println("Initialized test:Prompter")
	}
}

func WithShellExecutor(shellExecutor ports.ShellExecutor) func() {
	return func() {
		ShellExecutor = shellExecutor
		fmt.Println("Initialized test:ShellExecutor")
	}
}
