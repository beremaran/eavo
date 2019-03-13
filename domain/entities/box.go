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

//Fits returns true if box fits into given box
func (b *Box) Fits(o *Box) bool {
	rotations := []*Box{
		b,
		b.Rotate(types.AxisX),
		b.Rotate(types.AxisY),
		b.Rotate(types.AxisZ),
	}

	for i := 0; i < len(rotations); i++ {
		r := rotations[i]

		if r.Size.X <= o.Size.X &&
			r.Size.Y <= o.Size.Y &&
			r.Size.Z <= o.Size.Z {
			return true
		}
	}

	return false
}

//Rotate returns rotated version of Box
func (b *Box) Rotate(axis types.Axis) *Box {
	rotated := &Box{}
	err := copier.Copy(rotated, b)
	if err != nil {
		panic(err)
	}

	switch axis {
	case types.AxisX:
		rotated.Size.Y, rotated.Size.Z = rotated.Size.Z, rotated.Size.Y
		break
	case types.AxisY:
		rotated.Size.X, rotated.Size.Z = rotated.Size.Z, rotated.Size.X
		break
	case types.AxisZ:
		rotated.Size.X, rotated.Size.Y = rotated.Size.Y, rotated.Size.X
		break
	}

	return rotated
}
