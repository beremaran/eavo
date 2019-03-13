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

func TestBox_Fits(t *testing.T) {
	box1 := Box{
		Size: types.Vector3i{
			X: 10, Y: 10, Z: 10,
		},
	}

	box2 := Box{
		Size: types.Vector3i{
			X: 5, Y: 5, Z: 5,
		},
	}

	box3 := Box{
		Size: types.Vector3i{
			X: 10, Y: 20, Z: 10,
		},
	}

	if !box2.Fits(&box1) {
		t.Errorf("(5, 5, 5) actually fits (10, 10, 10) box, you know?")
	}

	if box3.Fits(&box1) {
		t.Errorf("(10, 20, 10) box can't be fitting into (10, 10, 10)")
	}
}

func TestBox_Rotate(t *testing.T) {
	box := Box{
		Size: types.Vector3i{
			X: 1, Y: 2, Z: 3,
		},
	}

	rX := box.Rotate(types.AxisX)
	rY := box.Rotate(types.AxisY)
	rZ := box.Rotate(types.AxisZ)

	if rX.Size.Y != 3 || rX.Size.Z != 2 {
		t.Errorf("x rotation didn't work")
	}

	if rY.Size.X != 3 || rY.Size.Z != 1 {
		t.Errorf("y rotation didn't work")
	}

	if rZ.Size.X != 2 || rZ.Size.Y != 1 {
		t.Errorf("z rotation didn't work")
	}
}
