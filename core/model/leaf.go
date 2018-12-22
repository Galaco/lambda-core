package model

import (
	"github.com/galaco/Gource-Engine/core/mesh"
)

// ClusterLeaf represents a single cluster that contains the contents of
// all the leafs that are contained withing it
type ClusterLeaf struct {
	Id          int16
	Faces       []mesh.Face
	StaticProps []*StaticProp
}
