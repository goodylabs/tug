package stageorchestrator_test

import (
	"errors"
	"testing"

	"github.com/goodylabs/tug/internal/utils/stageorchestrator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStageOrchestrator_ForwardOnly(t *testing.T) {
	steps := []stageorchestrator.StepFunc{
		func() (bool, error) { return true, nil },
		func() (bool, error) { return true, nil },
		func() (bool, error) { return true, nil },
	}

	o := stageorchestrator.NewStageOrchestrator(steps)

	err := o.Run()
	require.NoError(t, err)

	assert.Equal(t, len(steps), o.GetCurrentStep())
}

func TestStageOrchestrator_BackwardFromFirstStep_Exit(t *testing.T) {
	calls := 0
	steps := []stageorchestrator.StepFunc{
		func() (bool, error) {
			calls++
			return false, nil
		},
	}

	o := stageorchestrator.NewStageOrchestrator(steps)

	err := o.Run()
	require.NoError(t, err)

	assert.Equal(t, 0, o.GetCurrentStep())
	assert.Equal(t, 1, calls)
}

func TestStageOrchestrator_ErrorDuringStep(t *testing.T) {
	steps := []stageorchestrator.StepFunc{
		func() (bool, error) { return true, nil },
		func() (bool, error) { return false, errors.New("step failed") },
	}

	o := stageorchestrator.NewStageOrchestrator(steps)

	err := o.Run()
	require.Error(t, err)
	assert.EqualError(t, err, "step failed")
	assert.Equal(t, 1, o.GetCurrentStep())
}
