package aotree

import (
	"gitlab.com/beremaran/eavo/domain/entities"
	"testing"
)

func TestContainsNode(t *testing.T) {
	toSearch := NewAoNode(entities.Box{})
	nodes := []*AoNode{
		NewAoNode(entities.Box{}),
		toSearch,
		NewAoNode(entities.Box{}),
	}

	if !ContainsNode(nodes, toSearch) {
		t.Errorf("node not found")
	}
}
