package tree

import (
	"github.com/go-gl/mathgl/mgl32"
)

type Leaf struct {
	Id            int32
	FaceIndexList []uint16
	ClusterId     int16
	Min           mgl32.Vec3
	Max           mgl32.Vec3
	SkyVisible    bool
}

func (leaf *Leaf) IsLeaf() bool {
	return true
}
