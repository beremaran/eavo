package genetic

import (
	"gitlab.com/beremaran/eavo/domain/aotree"
	"gitlab.com/beremaran/eavo/domain/entities"
)

//FindUsedBoxes finds all boxes from given AO-tree which can be used for
//any box in problem context
func FindUsedBoxes(tree *aotree.AoNode, problem entities.Problem) map[*entities.Box]*entities.Box {
	problemBoxes := problem.Boxes
	storeNodes := tree.Root().StoreNodes()
	nodeMap := map[*entities.Box]*entities.Box{}

	for i := 0; i < len(storeNodes); i++ {
		node := storeNodes[i]

		for j := 0; j < len(problemBoxes); j++ {
			pBox := problemBoxes[j]
			if containsBoxKey(nodeMap, &pBox) {
				continue
			}

			if pBox.Fits(&node.Box) {
				nodeMap[&pBox] = &node.Box
				break
			}
		}
	}

	return nodeMap
}

func containsBoxKey(m map[*entities.Box]*entities.Box, b *entities.Box) bool {
	_, ok := m[b]
	return ok
}
