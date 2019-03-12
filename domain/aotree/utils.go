package aotree

import "gitlab.com/beremaran/eavo/domain/entities"

//ContainsNode searches for a node in given node slice
func ContainsNode(arr []*AoNode, n *AoNode) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == n {
			return true
		}
	}

	return false
}

//MapToBoxes maps a tree's store nodes to their bounding boxes
func MapToBoxes(node *AoNode) ([]*entities.Box, error) {
	var boxes []*entities.Box
	nodes := node.Root().StoreNodes()

	for i := 0; i < len(nodes); i++ {
		boxes = append(boxes, &nodes[i].Box)
	}

	return boxes, nil
}
