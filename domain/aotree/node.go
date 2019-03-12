package aotree

import (
	"errors"
	"fmt"
	"gitlab.com/beremaran/eavo/domain/entities"
	"gitlab.com/beremaran/eavo/domain/types"
)

const (
	//NodeCut
	NodeCut = 'c'
	//NodeStore
	NodeStore = 's'
)

var ErrAlreadyCut = errors.New("node is already cut")

//NodeType
type NodeType byte

//CutSpecs
type CutSpecs struct {
	Position int
	Axis     types.Axis
}

//AoNode
type AoNode struct {
	Parent   *AoNode
	Left     *AoNode
	Right    *AoNode
	Box      entities.Box
	CutSpecs CutSpecs
	NodeType NodeType
}

//NewAoNode
func NewAoNode(box entities.Box) *AoNode {
	return &AoNode{
		Parent:   nil,
		Left:     nil,
		Right:    nil,
		Box:      box,
		NodeType: NodeStore,
	}
}

//Uncut
func (n *AoNode) Uncut() {
	n.NodeType = NodeStore
	n.Left = nil
	n.Right = nil
}

//Cut
func (n *AoNode) Cut(position int, axis types.Axis) error {
	if n.NodeType == NodeCut || (n.Left != nil && n.Right != nil) {
		return ErrAlreadyCut
	}

	if position <= 0 || position >= n.Box.GetAxisLength(axis) {
		return fmt.Errorf("invalid cut position at %c=%d (%c=%d)", axis, position, axis, n.Box.GetAxisLength(axis))
	}

	n.CutSpecs.Position = position
	n.CutSpecs.Axis = axis
	n.NodeType = NodeCut

	n.Left = &AoNode{
		Parent: n,
	}

	n.Right = &AoNode{
		Parent: n,
	}

	n.Left.CalculateBox()
	n.Right.CalculateBox()
	return nil
}

//CalculateBox
func (n *AoNode) CalculateBox() {
	if n.Parent == nil {
		return
	}

	cutSpecs := n.Parent.CutSpecs
	n.Box = *n.Parent.Box.Clone()
	if n.Parent.Left == n {
		switch cutSpecs.Axis {
		case types.AxisX:
			n.Box.Size.X = cutSpecs.Position
			break
		case types.AxisY:
			n.Box.Size.Y = cutSpecs.Position
			break
		case types.AxisZ:
			n.Box.Size.Z = cutSpecs.Position
			break
		}
	} else {
		switch cutSpecs.Axis {
		case types.AxisX:
			n.Box.Size.X -= cutSpecs.Position
			n.Box.Position.X += cutSpecs.Position
			break
		case types.AxisY:
			n.Box.Size.Y -= cutSpecs.Position
			n.Box.Position.Y += cutSpecs.Position
			break
		case types.AxisZ:
			n.Box.Size.Z -= cutSpecs.Position
			n.Box.Position.Z += cutSpecs.Position
			break
		}
	}
}

//Root
func (n *AoNode) Root() *AoNode {
	node := n

	for node.Parent != nil {
		node = node.Parent
	}

	return node
}

//Traverse
func (n *AoNode) Traverse(filter func(node *AoNode) bool) []*AoNode {
	queue := []*AoNode{n.Root()}
	var visited []*AoNode

	for len(queue) > 0 {
		node := queue[0]
		if len(queue) >= 1 {
			queue = queue[1:]
		}

		if ContainsNode(visited, node) {
			continue
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}

		if filter == nil || filter(node) {
			visited = append(visited, node)
		}
	}

	return visited
}

//All
func (n *AoNode) All() []*AoNode {
	return n.Traverse(nil)
}

//StoreNodes
func (n *AoNode) StoreNodes() []*AoNode {
	return n.Traverse(func(node *AoNode) bool {
		return node.NodeType == NodeStore
	})
}

//CutNodes
func (n *AoNode) CutNodes() []*AoNode {
	return n.Traverse(func(node *AoNode) bool {
		return node.NodeType == NodeCut
	})
}
