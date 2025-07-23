package adapters

import (
	"github.com/goodylabs/tug/internal/adapters/dockercli"
	"github.com/goodylabs/tug/internal/adapters/prompter"
	"github.com/goodylabs/tug/internal/ports"
)

var DockerApi ports.DockerApi
var Prompter ports.Prompter

func InitializeDependencies(options ...func()) {
	if DockerApi == nil {
		DockerApi = dockercli.NewDockerApi()
	}
	if Prompter == nil {
		Prompter = prompter.NewPrompter()
	}

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
