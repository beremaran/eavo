package entities

import (
	"github.com/jinzhu/copier"
	"gitlab.com/beremaran/eavo/domain/types"
)

type Box struct {
	Size     types.Vector3i
	Position types.Vector3i
}

func (b *Box) GetAxisLength(axis types.Axis) int {
	switch axis {
	case types.AxisX:
		return b.Size.X
	case types.AxisY:
		return b.Size.Y
	case types.AxisZ:
		return b.Size.Z
	}

	return -1
}

func (b *Box) Weight() int {
	return b.Volume()
}

func (b *Box) Volume() int {
	return b.Size.X * b.Size.Y * b.Size.Z
}

func (b *Box) Clone() *Box {
	box := &Box{}
	err := copier.Copy(box, b)
	if err != nil {
		panic(err)
	}

	return box
}
