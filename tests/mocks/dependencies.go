package mocks

import (
	"github.com/goodylabs/tug/internal/app"
)

// func SetupPm2ManagerWithMocks(prompts []int, sshOutput string, sshErr error) ports.Pm2Manager {
// 	return pm2.NewPm2Manager(
// 		NewPrompterMock(prompts),
// 		NewSSHConnectorMock(sshOutput, sshErr),
// 	)
// }

func SetupConfigureUseCaseWithMocks(prompts []int) *app.ConfigureUseCase {
	return app.NewConfigureUseCase(
		NewPrompterMock(prompts),
	)
}
