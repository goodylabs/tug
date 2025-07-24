package application_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/tests/mocks"
	"github.com/goodylabs/tug/tests/testutils"
)

func init() {
	adapters.InitializeDependencies(
		adapters.WithShellExecutor(mocks.NewShellExecutor()),
	)
}

func TestUpdateTugUseCase(t *testing.T) {

	path, cleanup := testutils.CreateTestTugReleaseFile()
	defer cleanup()

	application.UpdateTugUseCase(path)

}
