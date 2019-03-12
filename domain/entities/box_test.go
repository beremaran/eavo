package entities

import (
	"gitlab.com/beremaran/eavo/domain/types"
	"testing"
)

func getBox(sizeX, sizeY, sizeZ int) *Box {
	return &Box{Size: types.Vector3i{X: sizeX, Y: sizeY, Z: sizeZ}}
}

func TestBox_GetAxisLength(t *testing.T) {
	box := getBox(10, 20, 30)

	checkAxisLength := func(expected int, axis types.Axis) {
		if box.GetAxisLength(axis) != expected {
			t.Errorf("%c axis length is not correct (%d != %d)", axis, expected, box.GetAxisLength(axis))
		}
	}

	checkAxisLength(10, types.AxisX)
	checkAxisLength(20, types.AxisY)
	checkAxisLength(30, types.AxisZ)
}

func TestBox_Volume(t *testing.T) {
	box := getBox(10, 20, 30)

	if box.Volume() != 10*20*30 {
		t.Errorf("wrong volume (%d != %d)", 10*20*30, box.Volume())
	}
}

func TestBox_Clone(t *testing.T) {
	box := getBox(10, 20, 30)
	clone := box.Clone()

	if box == clone {
		t.Errorf("clone returned same box")
	}

	if box.Size.X != clone.Size.X ||
		box.Size.Y != clone.Size.Y ||
		box.Size.Z != clone.Size.Z {
		t.Errorf("incorrect cloned box size")
	}
}
