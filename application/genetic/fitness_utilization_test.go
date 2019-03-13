package genetic

import (
	"gitlab.com/beremaran/eavo/domain/entities"
	"gitlab.com/beremaran/eavo/domain/genetic"
	"gitlab.com/beremaran/eavo/domain/types"
	"testing"
)

func TestUtilizationFitness_Calculate(t *testing.T) {
	container := entities.Container{
		Size: types.Vector3i{
			X: 10,
			Y: 10,
			Z: 10,
		},
	}

	context := genetic.Context{
		Problem: entities.Problem{
			Container: container,
			Boxes:     getRandomBoxes(5, container),
		},
	}

	fitness := UtilizationFitness{}
	fitness.SetContext(&context)
}
