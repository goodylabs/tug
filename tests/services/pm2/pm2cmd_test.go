package pm2_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/goodylabs/tug/internal/services/pm2"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSelectResourceOk(t *testing.T) {
	var pm2Manager *pm2.Pm2Manager
	var resource string
	var err error

	pm2Manager = mocks.SetupMockPm2Manager([]int{0}, output1, nil)
	resource, err = pm2Manager.SelectResource()
	assert.NoError(t, err)
	assert.Equal(t, "api-staging", resource)

	pm2Manager = mocks.SetupMockPm2Manager([]int{1}, output1, nil)
	resource, err = pm2Manager.SelectResource()
	assert.NoError(t, err)
	assert.Equal(t, "pm2-logrotate", resource)
}

func TestSelectResourceCommandError(t *testing.T) {
	pm2Manager := mocks.SetupMockPm2Manager([]int{0}, "", errors.New("command error"))
	resource, err := pm2Manager.SelectResource()
	assert.ErrorContains(t, err, "running PM2 jlist command: command error")
	assert.Equal(t, "", resource)
}

func TestSelectResourceInvalidOutput(t *testing.T) {
	pm2Manager := mocks.SetupMockPm2Manager([]int{0}, invalidOutput1, nil)
	resource, err := pm2Manager.SelectResource()
	assert.ErrorContains(t, err, "parsing PM2 list output")
	assert.Equal(t, "", resource)
}

func TestSelectCommandTemplateOk(t *testing.T) {
	testCases := []struct {
		prompterChoice int
		expected       string
	}{
		{
			prompterChoice: 1,
			expected:       "source ~/.nvm/nvm.sh; pm2 logs api-staging",
		},
		{
			prompterChoice: 4,
			expected:       "source ~/.nvm/nvm.sh; pm2 show api-staging",
		},
		{
			prompterChoice: 0,
			expected:       "source ~/.nvm/nvm.sh; pm2 describe api-staging",
		},
	}

	for _, testCase := range testCases {
		pm2Manager := mocks.SetupMockPm2Manager([]int{testCase.prompterChoice}, "", nil)
		templ, err := pm2Manager.SelectCommandTemplate()
		fullCmd := fmt.Sprintf(templ, "api-staging")
		assert.Equal(t, testCase.expected, fullCmd)
		assert.NoError(t, err)
	}
}
