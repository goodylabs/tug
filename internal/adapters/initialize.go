package adapters

import (
	"github.com/goodylabs/tug/internal/adapters/dockercli"
	"github.com/goodylabs/tug/internal/adapters/httpconnector"
	"github.com/goodylabs/tug/internal/adapters/prompter"
	"github.com/goodylabs/tug/internal/adapters/shellexecutor"
	"github.com/goodylabs/tug/internal/adapters/sshconnector"
	"github.com/goodylabs/tug/internal/ports"
)

var DockerApi ports.DockerApi
var Prompter ports.Prompter
var ShellExecutor ports.ShellExecutor
var HttpConnector ports.HttpConnector
var SSHConnector ports.SSHConnector

func InitializeDependencies(options ...func()) {
	DockerApi = dockercli.NewDockerApi()
	Prompter = prompter.NewPrompter()
	ShellExecutor = shellexecutor.NewShellExecutor()
	HttpConnector = httpconnector.NewHttpConnector()
	SSHConnector = sshconnector.NewSSHConnector()

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

func WithSSHConnector(sshConnector ports.SSHConnector) func() {
	return func() {
		SSHConnector = sshConnector
	}
}
