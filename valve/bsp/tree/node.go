package tree

import "github.com/go-gl/mathgl/mgl32"

type INode interface {
	IsLeaf() bool
}

type Node struct {
	Children [2]INode
	Min mgl32.Vec3
	Max mgl32.Vec3
}

func (node *Node) IsLeaf() bool {
	return false
}
func (node *Node) AddChild(index int, child INode) {
	node.Children[index] = child
}