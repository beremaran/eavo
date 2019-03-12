package entities

import "gitlab.com/beremaran/eavo/domain/types"

//Container encapsulates the properties of a Container entity
type Container struct {
	Size types.Vector3i
}

//Volume returns volume of the container
func (c *Container) Volume() int {
	return c.Size.X * c.Size.Y * c.Size.Z
}
