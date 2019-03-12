package entities

import (
	"beremaran/eavo/domain/types"
	"github.com/jinzhu/copier"
)

type Box struct {
	Size     types.Vector3i
	Position types.Vector3i
}

func (b *Box) Weight() int {
	return b.Volume()
}

func (b *Box) Volume() int {
	return b.Size.X * b.Size.Y * b.Size.Z
}

func (b *Box) Clone() Box {
	box := Box{}
	err := copier.Copy(&box, b)
	if err != nil {
		panic(err)
	}

	return box
}
