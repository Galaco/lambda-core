package tree

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/go-gl/mathgl/mgl32"
)

type Leaf struct {
	Id            int32
	FaceIndexList []uint16
	ClusterId     int16
	Min           mgl32.Vec3
	Max           mgl32.Vec3
	Faces		  []interfaces.IPrimitive
}

func (leaf *Leaf) IsLeaf() bool {
	return true
}
