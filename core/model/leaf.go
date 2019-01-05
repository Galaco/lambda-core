package model

import (
	"github.com/galaco/Gource-Engine/core/mesh"
	"github.com/go-gl/mathgl/mgl32"
)

// ClusterLeaf represents a single cluster that contains the contents of
// all the leafs that are contained withing it
type ClusterLeaf struct {
	Id          int16
	Faces       []mesh.Face
	StaticProps []*StaticProp
	DispFaces   []int
	Mins 		mgl32.Vec3
	Maxs		mgl32.Vec3
}
