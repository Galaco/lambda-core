package vis

import "github.com/galaco/Gource/valve/vis/tree"

type Cache struct {
	Leaves     []*tree.Leaf
	ClusterId  int16
	SkyVisible bool
}
