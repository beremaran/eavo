package entities

import (
	"beremaran/eavo/domain/types"
	"testing"
)

func getContainer(sizeX, sizeY, sizeZ int) Container {
	return Container{Size: types.Vector3i{X: sizeX, Y: sizeY, Z: sizeZ}}
}

func TestContainer_Volume(t *testing.T) {
	container := getContainer(10, 20, 30)

	if container.Volume() != 10*20*30 {
		t.Errorf("incorrect container volume (%d != %d)", 10*20*30, container.Volume())
	}
}
