package mocks

import (
	"github.com/goodylabs/tug/internal/services/pm2"
	"github.com/goodylabs/tug/internal/services/tughelper"
)

func SetupPm2ManagerWithMocks(prompts []int, sshOutput string, sshErr error) *pm2.Pm2Manager {
	return pm2.NewPm2Manager(
		NewPrompterMock(prompts),
		NewSSHConnectorMock(sshOutput, sshErr),
	)
}

func SetupTugHelperWithMocks(prompts []int) *tughelper.TugHelper {
	return tughelper.NewTugHelper(
		NewPrompterMock(prompts),
	)
}
