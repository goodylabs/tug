package services

import (
	"github.com/goodylabs/tug/internal/adapters"
	"github.com/goodylabs/tug/internal/constants"
	"github.com/goodylabs/tug/internal/dto"
)

func ChoseContainer(containers []dto.ContainerDTO) dto.ContainerDTO {
	var names []string
	for _, container := range containers {
		names = append(names, container.Name)
	}
	selectedName := adapters.Prompter.ChooseFromList(names, "Chose container")
	for _, container := range containers {
		if selectedName == container.Name {
			return container
		}
	}
	panic(constants.PANIC)
}
