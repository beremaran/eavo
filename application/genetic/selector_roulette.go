package genetic

import (
	"errors"
	"gitlab.com/beremaran/eavo/domain/genetic"
	"math/rand"
)

//RouletteWheelSelector selects an individual from given population
//implying Roulette Wheel Selection method.
type RouletteWheelSelector struct {
}

//Select selects an individual from given population
func (s *RouletteWheelSelector) Select(population []*genetic.Individual) (*genetic.Individual, error) {
	totalFitness := getTotalFitness(population)
	if totalFitness == 0 {
		return getRandomIndividual(population)
	}

	for i := 0; i < len(population); i++ {
		if rand.Float64() <= (population[i].FitnessScore / totalFitness) {
			return population[i], nil
		}
	}

	return getRandomIndividual(population)
}

func getTotalFitness(population []*genetic.Individual) float64 {
	totalFitness := 0.0

	for i := 0; i < len(population); i++ {
		totalFitness += population[i].FitnessScore
	}

	return totalFitness
}

func getRandomIndividual(population []*genetic.Individual) (*genetic.Individual, error) {
	if len(population) == 0 {
		return nil, errors.New("empty population")
	}

	return population[rand.Intn(len(population))], nil
}
