package ports

import "github.com/goodylabs/tug/internal/dto"

type Prompter interface {
	ChooseFromList([]string, string) string
}

type DockerApi interface {
	ConfigureDocker(targetIP string)
	ListServices() []dto.ServiceDTO
	ListContainers() []dto.ContainerDTO
	ExecLogsLive(containerId string)
	ExecShell(containerId string)
	ExecDjangoShell(containerId string)
	ExecLogs(containerId string)
}

type ShellExecutor interface {
	Exec(command string) error
}

type HttpConnector interface {
	HttpGetReq(url string, target any) error
}

type SSHConnector interface {
	OpenConnection(user, host string, port int) error
	CloseConnection() error
	RunCommand(cmd string) (string, error)
}
