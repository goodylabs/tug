package stageorchestrator

import "fmt"

type StepFunc func() (bool, error)

type StageOrchestrator struct {
	steps       []StepFunc
	currentStep int
}

func NewStageOrchestrator(steps []StepFunc) *StageOrchestrator {
	return &StageOrchestrator{
		steps: steps,
	}
}

func (s *StageOrchestrator) Run() error {
	for s.currentStep < len(s.steps) {
		nextStep, err := s.steps[s.currentStep]()
		if err != nil {
			return err
		}
		if nextStep {
			s.currentStep++
			continue
		}
		if s.currentStep == 0 {
			fmt.Println("Exiting...")
			return nil
		}
		s.currentStep--
	}
	return nil
}

func (s *StageOrchestrator) GetCurrentStep() int {
	return s.currentStep
}
