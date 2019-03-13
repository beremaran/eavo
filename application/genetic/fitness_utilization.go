package genetic

import (
	"gitlab.com/beremaran/eavo/domain/aotree"
	"gitlab.com/beremaran/eavo/domain/entities"
	"gitlab.com/beremaran/eavo/domain/genetic"
)

//UtilizationFitness calculates a fitness score considering
//volume utilization of container
type UtilizationFitness struct {
	context *genetic.Context
}

//GetWeight returns corresponding fitness score weight
func (f *UtilizationFitness) GetWeight() float64 {
	return 1.0
}

//SetContext sets genetic context for the fitness function
func (f *UtilizationFitness) SetContext(ctx *genetic.Context) {
	f.context = ctx
}

//Calculate calculates a fitness score considering
//volume utilization of container
func (f *UtilizationFitness) Calculate(individual *genetic.Individual) (float64, error) {
	containerVolume := float64(f.context.Problem.Container.Volume())
	totalBoxVolume := 0.0

	problemBoxes := f.context.Problem.Boxes
	nodeMap := map[*entities.Box]*aotree.AoNode{}
	storeNodes := individual.Genome.Root().StoreNodes()
	for i := 0; i < len(storeNodes); i++ {
		node := storeNodes[i]

		for j := 0; j < len(problemBoxes); j++ {
			pBox := problemBoxes[j]
			if containsBoxKey(nodeMap, &pBox) {
				continue
			}

			if pBox.Fits(&node.Box) {
				nodeMap[&pBox] = node
				break
			}
		}
	}

	for pBox := range nodeMap {
		totalBoxVolume += float64(pBox.Volume())
	}

	return totalBoxVolume / containerVolume, nil
}

func containsBoxKey(m map[*entities.Box]*aotree.AoNode, b *entities.Box) bool {
	_, ok := m[b]
	return ok
}
