package genetic

import "gitlab.com/beremaran/eavo/domain/entities"

//Context encapsulates the context of a genetic optimization process
type Context struct {
	Problem                entities.Problem
	MutationRate           float64
	CrossoverRate          float64
	PopulationSize         int
	MaxGenerations         int
	MaxRepeatedGenerations int
	NumberOfElites         int
}
