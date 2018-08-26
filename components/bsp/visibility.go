package bsp

import (
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
)

var currentLeaf *tree.Leaf

func RebuildVisibilityList(treeList []tree.Node, position mgl32.Vec3) ([]uint16, *tree.Leaf) {
	visibleList := []uint16{}
	for _, root := range treeList {
		depth := 0
		visibleList = append(visibleList, rebuildVisibilityListRecursive(&root, position, depth)...)
	}
	return visibleList, currentLeaf
}

func rebuildVisibilityListRecursive(node tree.INode, position mgl32.Vec3, depth int) []uint16 {
	ret := []uint16{}
	if node.IsLeaf() == false {
		for _,child := range node.(*tree.Node).Children {
			ret = append(ret, rebuildVisibilityListRecursive(child, position, depth + 1)...)
		}
		return ret
	} else {

		if depth > 0 {
			// Skip this node and its children if we aren't in it
			if IsPointInLeaf(position, node.(*tree.Leaf).Min, node.(*tree.Leaf).Max) {
				currentLeaf = node.(*tree.Leaf)
			}
		}
		return node.(*tree.Leaf).FaceIndexList
	}

	return ret
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
