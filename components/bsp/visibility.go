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

		if IsPointInFrontOfPlane(position, localNode.Plane) == true {
			return findCurrentLeafRecursive(localNode.Children[0], position)
		} else {
			return findCurrentLeafRecursive(localNode.Children[1], position)
		}
	}

	return node.(*tree.Leaf)
}

func IsPointInFrontOfPlane(point mgl32.Vec3, nodePlane *plane.Plane) bool {
	// vector from origin to plane || is also coordinates on place
	planeToOrigin := nodePlane.Normal.Mul(nodePlane.Distance)
	planeToPoint := point.Sub(planeToOrigin)
	dot := nodePlane.Normal.Dot(planeToPoint.Normalize())
	return dot > 0
}

func BuildFaceListForVisibleClusters(nodeTree []tree.Node, clusterList []int16) []uint16 {
	return recursiveBuildFaceIndexList(&nodeTree[0], []uint16{}, clusterList)
}

func recursiveBuildFaceIndexList(node tree.INode, faceList []uint16, clusterList []int16) []uint16 {
	if node.IsLeaf() {
		for _, v := range clusterList {
			if v == node.(*tree.Leaf).ClusterId {
				faceList = append(faceList, node.(*tree.Leaf).FaceIndexList...)
				break
			}
		}
	} else {
		for _, child := range node.(*tree.Node).Children {
			faceList = recursiveBuildFaceIndexList(child, faceList, clusterList)
		}
	}

	return faceList
}
