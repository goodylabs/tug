package app_test

import (
	"os"
	"testing"

	"github.com/goodylabs/tug/internal/tughelper"
	"github.com/goodylabs/tug/pkg/config"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInitializeUseCaseOk(t *testing.T) {
	testCases := []struct {
		promptChoices []int
		expected      string
	}{
		{
			promptChoices: []int{1},
			expected:      "some_priv_key",
		},
		{
			promptChoices: []int{0},
			expected:      "id_rsa",
		},
	}

	for _, test := range testCases {
		os.Remove(config.GetTugConfigPath())

		useCase := mocks.SetupInitializeUseCaseWithMocks(test.promptChoices)
		err := useCase.Execute()
		assert.NoError(t, err)

		tugConfig, err := tughelper.GetTugConfig()
		assert.NoError(t, err, "Expected no error, got: %v", err)
		assert.Contains(t, tugConfig.SSHKeyPath, test.expected)
	}
}
