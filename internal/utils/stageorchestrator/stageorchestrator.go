package stageorchestrator

import "fmt"

type StepFunc func() (bool, error)

type StageOrchestrator struct {
	steps []StepFunc
}

func NewStageOrchestrator(steps []StepFunc) *StageOrchestrator {
	return &StageOrchestrator{
		steps: steps,
	}
}

func (s *StageOrchestrator) Run() error {
	currentStep := 0

	for currentStep < len(s.steps) {
		nextStep, err := s.steps[currentStep]()
		if err != nil {
			return err
		}
		if nextStep {
			currentStep++
			continue
		}
		if currentStep == 0 {
			fmt.Println("Exiting...")
			return nil
		}
		currentStep--
	}
	return nil
}
