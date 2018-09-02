package tree

import (
	"github.com/galaco/bsp/primitives/plane"
	"github.com/go-gl/mathgl/mgl32"
)

type INode interface {
	IsLeaf() bool
}

type Node struct {
	Id       int32
	Children [2]INode
	Min      mgl32.Vec3
	Max      mgl32.Vec3
	Plane    *plane.Plane
}

func (node *Node) IsLeaf() bool {
	return false
}
func (node *Node) AddChild(index int, child INode) {
	node.Children[index] = child
}
