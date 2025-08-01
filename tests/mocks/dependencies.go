package mocks

import (
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/internal/modules/pm2"
	"github.com/goodylabs/tug/internal/ports"
)

func SetupPm2ManagerWithMocks(prompts []int, sshOutput string, sshErr error) ports.Pm2Manager {
	return pm2.NewPm2Manager(
		NewPrompterMock(prompts),
		NewSSHConnectorMock(sshOutput, sshErr),
	)
}

func SetupInitializeUseCaseWithMocks(prompts []int) *application.InitializeUseCase {
	return application.NewInitializeUseCase(
		NewPrompterMock(prompts),
	)
}
