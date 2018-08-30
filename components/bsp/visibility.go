package bsp

import (
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
)

var currentLeaf *tree.Leaf

func FindCurrentLeaf(treeList []tree.Node, position mgl32.Vec3) *tree.Leaf {
	currentLeaf = nil
	for _, root := range treeList {
		findCurrentLeafRecursive(&root, position)
	}
	return currentLeaf
}

func findCurrentLeafRecursive(node tree.INode, position mgl32.Vec3) {
	if node.IsLeaf() == false {
		for _, child := range node.(*tree.Node).Children {
			findCurrentLeafRecursive(child, position)
		}
	} else {
		// Skip this node and its children if we aren't in it
		if IsPointInLeaf(position, node.(*tree.Leaf).Min, node.(*tree.Leaf).Max) {
			currentLeaf = node.(*tree.Leaf)
		}
	}
}

func IsPointInLeaf(point mgl32.Vec3, min mgl32.Vec3, max mgl32.Vec3) bool {
	if point.X() < min.X() ||
		point.X() > max.X() ||
		point.Y() < min.Y() ||
		point.Y() > max.Y() ||
		point.Z() < min.Z() ||
		point.Z() > max.Z() {
		return false
	}
	return true
}

func BuildFaceListForVisibleClusters(nodeTree []tree.Node, clusterList []int16) []uint16 {
	faceList := []uint16{}
	for _, root := range nodeTree {
		faceList = append(faceList, recursiveBuildFaceIndexList(&root, faceList, clusterList)...)
	}

	return faceList
}

func recursiveBuildFaceIndexList(node tree.INode, faceList []uint16, clusterList []int16) []uint16 {
	if node.IsLeaf() {
		clusterId := node.(*tree.Leaf).ClusterId
		for _, v := range clusterList {
			if v == clusterId {
				return node.(*tree.Leaf).FaceIndexList
			}
		}
		return []uint16{}
	} else {
		for _, child := range node.(*tree.Node).Children {
			faceList = recursiveBuildFaceIndexList(child, faceList, clusterList)
		}
	}

	return faceList
}
