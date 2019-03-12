package entities

import (
	"github.com/jinzhu/copier"
	"gitlab.com/beremaran/eavo/domain/types"
)

//Box encapsulates properties of a Box entity
type Box struct {
	Size     types.Vector3i
	Position types.Vector3i
}

//GetAxisLength returns the length of a given axis
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

//Weight returns weight of the box
func (b *Box) Weight() int {
	return b.Volume()
}

//Volume returns the volume of the box
func (b *Box) Volume() int {
	return b.Size.X * b.Size.Y * b.Size.Z
}

//Clone creates a independent copy (deep-copy) from box
func (b *Box) Clone() *Box {
	box := &Box{}
	err := copier.Copy(box, b)
	if err != nil {
		panic(err)
	}

	return box
}
