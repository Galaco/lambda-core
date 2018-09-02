package bsp

import (
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
)

func FindCurrentLeaf(treeList []tree.Node, position mgl32.Vec3) *tree.Leaf {
	// offset [0] is always worldspawn
	// see: https://developer.valvesoftware.com/wiki/Source_BSP_File_Format#Model
	return findCurrentLeafRecursive(&treeList[0], position)
}

func findCurrentLeafRecursive(node tree.INode, position mgl32.Vec3) *tree.Leaf {
	// treat as a npde
	if node.IsLeaf() != true {
		localNode := node.(*tree.Node)

		if isPointInFrontOfPlane(position, localNode.Plane) == true {
			return findCurrentLeafRecursive(localNode.Children[0], position)
		} else {
			return findCurrentLeafRecursive(localNode.Children[1], position)
		}
	}

	return node.(*tree.Leaf)
}

// Check if viewpoint is in front or behind the split plane
// dot product of place to origin & plane to viewpoint
func isPointInFrontOfPlane(point mgl32.Vec3, nodePlane *plane.Plane) bool {
	dist := (nodePlane.Normal.X()*point.X() +
		nodePlane.Normal.Y()*point.Y() +
		nodePlane.Normal.Z()*point.Z()) - nodePlane.Distance

	return dist >= 0
	//
	//planeToOrigin := nodePlane.Normal.Mul(nodePlane.Distance)
	//planeToPoint := point.Sub(planeToOrigin)
	//
	//return nodePlane.Normal.Dot(planeToPoint.Normalize()) > 0
}
