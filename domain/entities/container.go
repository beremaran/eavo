package entities

import "gitlab.com/beremaran/eavo/domain/types"

type Container struct {
	Size types.Vector3i
}

func (c *Container) Volume() int {
	return c.Size.X * c.Size.Y * c.Size.Z
}
