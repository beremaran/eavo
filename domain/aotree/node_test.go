package aotree

import (
	"gitlab.com/beremaran/eavo/domain/entities"
	"gitlab.com/beremaran/eavo/domain/types"
	"testing"
)

func get3LevelTree() *AoNode {
	root := NewAoNode(entities.Box{Size: types.Vector3i{X: 10, Y: 20, Z: 30}})

	_ = root.Cut(3, types.AxisX)
	_ = root.Left.Cut(5, types.AxisY)
	_ = root.Right.Cut(5, types.AxisZ)
	_ = root.Left.Left.Cut(1, types.AxisZ)
	_ = root.Right.Right.Cut(1, types.AxisX)

	return root
}

func TestNewAoNode(t *testing.T) {
	node := NewAoNode(entities.Box{})

	if node == nil {
		t.Errorf("could not create a new node")
	}
}

func TestAoNode_Root(t *testing.T) {
	nodeRoot := NewAoNode(entities.Box{})
	nodeLeft := NewAoNode(entities.Box{})
	nodeLeftRight := NewAoNode(entities.Box{})

	nodeRoot.Left = nodeLeft
	nodeRoot.Left.Parent = nodeRoot

	nodeLeft.Right = nodeLeftRight
	nodeLeft.Right.Parent = nodeLeft

	root := nodeLeftRight.Root()
	if root != nodeRoot {
		t.Errorf("wrong node returned as root")
	}
}

func TestAoNode_Cut(t *testing.T) {
	root := NewAoNode(entities.Box{Size: types.Vector3i{X: 10, Y: 20, Z: 30}})
	err := root.Cut(3, types.AxisX)
	if err != nil {
		t.Error(err)
	}

	if root.Left == nil {
		t.Errorf("left child is nil")
	}

	if root.Right == nil {
		t.Errorf("right child is nil")
	}

	if root.NodeType != NodeCut {
		t.Errorf("node type is incorrect (%c != %c)", root.NodeType, NodeCut)
	}

	if root.Left.Parent != root {
		t.Errorf("left child parent is wrong")
	}

	if root.Right.Parent != root {
		t.Errorf("right child parent is wrong")
	}

	if root.Left.Box.Size.X != 3 {
		t.Errorf("left child box size is wrong (%d != %d)", root.Left.Box.Size.X, 3)
	}

	if root.Right.Box.Size.X != 7 {
		t.Errorf("right child box size is wrong (%d != %d)", root.Right.Box.Size.X, 7)
	}

	if root.Left.Box.Position.X != 0 {
		t.Errorf("left child box position is wrong (%d != %d)", root.Left.Box.Position.X, 0)
	}

	if root.Right.Box.Position.X != 3 {
		t.Errorf("right child box position is wrong (%d != %d)", root.Right.Box.Position.X, 3)
	}
}

func TestAoNode_All(t *testing.T) {
	tree := get3LevelTree()
	nodes := tree.All()

	if len(nodes) != 11 {
		t.Errorf("wrong number of nodes in tree (%d != %d)", len(nodes), 11)
	}
}

func TestAoNode_CutNodes(t *testing.T) {
	tree := get3LevelTree()
	nodes := tree.CutNodes()

	if len(nodes) != 5 {
		t.Errorf("wrong number of cut nodes in tree (%d != %d)", len(nodes), 5)
	}
}

func TestAoNode_StoreNodes(t *testing.T) {
	tree := get3LevelTree()
	nodes := tree.StoreNodes()

	if len(nodes) != 6 {
		t.Errorf("wrong number of store nodes in tree (%d != %d)", len(nodes), 6)
	}
}
