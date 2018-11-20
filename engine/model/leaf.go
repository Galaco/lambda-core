package model

import (
	"github.com/galaco/Gource-Engine/engine/mesh"
)

type ClusterLeaf struct {
	Id          int16
	Faces       []mesh.Face
	StaticProps []*StaticProp
}