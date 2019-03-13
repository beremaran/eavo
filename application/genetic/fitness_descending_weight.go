package genetic

import "gitlab.com/beremaran/eavo/domain/genetic"

//DescendingWeightFitness directs local search to have boxes
//have descending weights when stacking up
type DescendingWeightFitness struct {
	context *genetic.Context
}

//GetWeight returns weight of DescendingWeightFitness
func (*DescendingWeightFitness) GetWeight() float64 {
	return 1.0
}

//SetContext sets genetic context for the fitness function
func (f *DescendingWeightFitness) SetContext(ctx *genetic.Context) {
	f.context = ctx
}

//Calculate directs local search to have boxes
//have descending weights when stacking up
func (f *DescendingWeightFitness) Calculate(individual *genetic.Individual) (float64, error) {
	weightedWeights := 0.0
	boxMap := FindUsedBoxes(individual.Genome, f.context.Problem)

	for pBox, tBox := range boxMap {
		weightedWeights += float64(tBox.Position.Y * pBox.Weight())
	}

	return 1.0 / (1.0 + weightedWeights), nil
}
