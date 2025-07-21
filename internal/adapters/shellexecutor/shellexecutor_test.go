package shellexecutor_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters/shellexecutor"
	"github.com/goodylabs/tug/internal/ports"
	"github.com/stretchr/testify/assert"
)

var executor ports.ShellExecutor

func init() {
	executor = shellexecutor.NewShellExecutor()
}

func TestShellExecBasic(t *testing.T) {
	assert.NoError(t, executor.Exec("echo ok"))

	assert.Error(t, executor.Exec("false"))
}
