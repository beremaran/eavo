package entities

import "beremaran/eavo/domain/types"

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
