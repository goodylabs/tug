package tughelper_test

import (
	"path/filepath"
	"testing"

	"github.com/goodylabs/tug/internal/config"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/stretchr/testify/assert"
)

type tugHelperTestCase struct {
	promptChoices []int
	expected      string
	value         string
}

func TestGetSSHDirPathOK(t *testing.T) {
	testCases := []tugHelperTestCase{
		{
			promptChoices: []int{0},
			expected:      "id_rsa",
		},
		{
			promptChoices: []int{1},
			expected:      "some_priv_key",
		},
	}

	for _, testCase := range testCases {
		tugHelper := mocks.SetupTugHelperWithMocks(testCase.promptChoices)
		sshDirPath := filepath.Join(config.BASE_DIR, ".ssh")
		result, err := tugHelper.GetSSHDirPath(sshDirPath)
		assert.Equal(t, testCase.expected, result)
		assert.NoError(t, err, "Expected no error, got: %v", err)
	}
}
