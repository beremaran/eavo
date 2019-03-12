package genetic

import (
	"github.com/jinzhu/copier"
	"gitlab.com/beremaran/eavo/domain/aotree"
)

//Individual is core element of the genetic algorithm. It's what is processed
//during the optimization
type Individual struct {
	FitnessScore float64
	Genome       *aotree.AoNode
}

//Copy makes a deep-copy of an Individual
func (i *Individual) Copy() *Individual {
	individual := &Individual{}
	err := copier.Copy(individual, i)
	if err != nil {
		panic(err)
	}

	return individual
}
