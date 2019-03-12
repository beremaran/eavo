package genetic

import (
	"gitlab.com/beremaran/eavo/domain/aotree"
	"gitlab.com/beremaran/eavo/domain/genetic"
	"gitlab.com/beremaran/eavo/domain/types"
	"math"
	"math/rand"
)

//Generator generates a new Individual
type Generator struct {
	context *genetic.Context
}

//SetContext sets genetic context of the generator
func (g *Generator) SetContext(ctx *genetic.Context) {
	g.context = ctx
}

//Generate generates a new Individual
func (g *Generator) Generate() (*genetic.Individual, error) {
	container := g.context.Problem.Container
	node := aotree.NewAoNode(container.ToBox())

	nCuts := rand.Intn(int(math.Log2(float64(len(g.context.Problem.Boxes))))) + 1
	for nCuts > 0 {
		randomNode := randomStoreNode(node)

		axis := getRandomAxis()
		if randomNode.Box.GetAxisLength(axis) > 1 {
			position := rand.Intn(randomNode.Box.GetAxisLength(axis)-1) + 1
			_ = randomNode.Cut(position, axis)
		}

		nCuts--
	}

	return &genetic.Individual{Genome: node}, nil
}

func getRandomAxis() types.Axis {
	return []types.Axis{types.AxisX, types.AxisY, types.AxisZ}[rand.Intn(3)]
}

func randomStoreNode(node *aotree.AoNode) *aotree.AoNode {
	storeNodes := node.Root().StoreNodes()
	return storeNodes[rand.Intn(len(storeNodes))]
}
