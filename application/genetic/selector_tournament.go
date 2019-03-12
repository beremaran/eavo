package genetic

import (
	"gitlab.com/beremaran/eavo/domain/genetic"
	"math/rand"
)

//TournamentSelector selects an individual implying Tournament Selection
type TournamentSelector struct {
	TournamentSize int
}

//Select selects from an individual from given population
func (s *TournamentSelector) Select(population []*genetic.Individual) (*genetic.Individual, error) {
	return s.tournament(population, 1)
}

func (s *TournamentSelector) tournament(population []*genetic.Individual, level int) (*genetic.Individual, error) {
	if level == s.TournamentSize {
		return population[rand.Intn(len(population))], nil
	}

	left, err := s.tournament(population, level+1)
	if err != nil {
		return nil, err
	}

	right, err := s.tournament(population, level+1)
	if err != nil {
		return nil, err
	}

	if left.FitnessScore > right.FitnessScore {
		return left, nil
	}

	return right, nil
}
