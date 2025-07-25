package ports

import "github.com/goodylabs/tug/internal/dto"

type Pm2Manager interface {
	LoadPm2Config(string, *dto.EconsystemConfigDTO) error
	JsonOutputHandler(string, any) error
}

type DockerManager interface {
	GetTargetIp(scriptAbsPath string) (string, error)
	ListContainers() []dto.ContainerDTO
	ChoseContainer(containers []dto.ContainerDTO) dto.ContainerDTO
	SelectAndExecuteCommand(container dto.ContainerDTO)
}
