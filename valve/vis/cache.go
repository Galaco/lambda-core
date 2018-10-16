package vis

import (
	"github.com/galaco/bsp/primitives/leaf"
)

type Cache struct {
	Leaves     []*leaf.Leaf
	ClusterId  int16
	SkyVisible bool
	Faces	   []uint16
}
