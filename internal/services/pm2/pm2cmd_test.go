package pm2_test

import (
	"errors"
	"testing"

	"github.com/goodylabs/tug/internal/services/pm2"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetPm2ProcessesOk(t *testing.T) {
	var pm2Manager *pm2.Pm2Manager
	var resources []string
	var err error

	pm2Manager = mocks.SetupPm2ManagerWithMocks([]int{0}, output1, nil)
	resources, err = pm2Manager.GetPm2Processes()
	assert.NoError(t, err)
	assert.Equal(t, "api-staging", resources)

	pm2Manager = mocks.SetupPm2ManagerWithMocks([]int{1}, output1, nil)
	resources, err = pm2Manager.GetPm2Processes()
	assert.NoError(t, err)
	assert.Equal(t, "pm2-logrotate", resources)
}

func TestGetPm2ProcessesCommandError(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, "", errors.New("command error"))
	resource, err := pm2Manager.GetPm2Processes()
	assert.ErrorContains(t, err, "running PM2 jlist command: command error")
	assert.Equal(t, "", resource)
}

func TestGetPm2ProcessesInvalidOutput(t *testing.T) {
	pm2Manager := mocks.SetupPm2ManagerWithMocks([]int{0}, invalidOutput1, nil)
	resource, err := pm2Manager.GetPm2Processes()
	assert.ErrorContains(t, err, "parsing PM2 list output")
	assert.Equal(t, "", resource)
}
