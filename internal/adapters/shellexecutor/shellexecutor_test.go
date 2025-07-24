package shellexecutor_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters/shellexecutor"
	"github.com/stretchr/testify/assert"
)

func TestShellExecBasic(t *testing.T) {
	executor := shellexecutor.NewShellExecutor()
	assert.NoError(t, executor.Exec("echo ok"))

	assert.Error(t, executor.Exec("false"))
}
