package services

import (
	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/dto"
)

const (
	cmdLogsLive    string = "Logs (live)"
	cmdShell       string = "Shell"
	cmdDjangoShell string = "Django shell"
	cmdLogs        string = "Logs"
)

var AllCommandTypes = []string{
	string(cmdLogsLive),
	string(cmdShell),
	string(cmdDjangoShell),
	string(cmdLogs),
}

func SelectAndExecCommandOnDockerContainer(container dto.ContainerDTO) {
	choice := adapters.Prompter.ChooseFromList(AllCommandTypes, "What do you want?")

	switch choice {
	case cmdLogsLive:
		adapters.DockerApi.ExecLogsLive(container.ID)
	case cmdShell:
		adapters.DockerApi.ExecShell(container.ID)
	case cmdDjangoShell:
		adapters.DockerApi.ExecDjangoShell(container.ID)
	case cmdLogs:
		adapters.DockerApi.ExecLogs(container.ID)
	}
}
