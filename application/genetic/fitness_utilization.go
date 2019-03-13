package genetic

import (
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

	nodeMap := FindUsedBoxes(individual.Genome, f.context.Problem)
	for pBox := range nodeMap {
		totalBoxVolume += float64(pBox.Volume())
	}

	return totalBoxVolume / containerVolume, nil
}
