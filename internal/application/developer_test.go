package application_test

import (
	"testing"

	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/application"
	"github.com/goodylabs/tug/tests/mocks"
)

func TestDeveloperForLocal(t *testing.T) {

	adapters.InitializeDependencies(
		adapters.WithPrompter(mocks.NewPrompterMock([]int{0, 3})),
	)

	options := application.DeveloperOptions{
		EnvDir: "localhost",
	}

	application.DeveloperUseCase(&options)

}
