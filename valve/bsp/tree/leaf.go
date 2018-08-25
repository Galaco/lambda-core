package tree

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
)

type Leaf struct {
	FaceIndexList []uint16
	FaceList []interfaces.IPrimitive
}

func (leaf *Leaf) IsLeaf() bool {
	return true
}

func (leaf *Leaf) AddFace(face interfaces.IPrimitive) {
	leaf.FaceList = append(leaf.FaceList, face)
}