package genetic

import (
	"gitlab.com/beremaran/eavo/domain/aotree"
	"gitlab.com/beremaran/eavo/domain/entities"
	"gitlab.com/beremaran/eavo/domain/genetic"
	"gitlab.com/beremaran/eavo/domain/types"
	"math/rand"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
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

	generator := Generator{}
	generator.SetContext(&context)
	var individuals []*genetic.Individual

	for i := 0; i < 100; i++ {
		individual, err := generator.Generate()
		if err != nil {
			t.Error(err)
		}

		individuals = append(individuals, individual)
	}

	for i := 0; i < 100; i++ {
		tree := individuals[i].Genome

		if tree.Root().NodeType == aotree.NodeStore {
			t.Errorf("1-length tree is found")
		}
	}
}

func getRandomBoxes(n int, container entities.Container) []entities.Box {
	var boxes []entities.Box

	for n > 0 {
		boxes = append(boxes, entities.Box{
			Size: types.Vector3i{
				X: rand.Intn(container.Size.X) + 1,
				Y: rand.Intn(container.Size.Y) + 1,
				Z: rand.Intn(container.Size.Z) + 1,
			},
		})

		n--
	}

	return boxes
}
