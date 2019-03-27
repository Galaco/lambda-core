package model

import (
	"github.com/galaco/Lambda-Core/core/mesh"
)

// ClusterLeaf represents a single cluster that contains the contents of
// all the leafs that are contained within it
type ClusterLeaf struct {
	Id          int16
	Faces       []mesh.Face
	StaticProps []*StaticProp
	DispFaces   []int
}
