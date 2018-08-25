package bsp

import (
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/galaco/go-me-engine/engine/interfaces"
)

func RebuildVisibilityList(treeList []tree.Node, position mgl32.Vec3) (visibleList []interfaces.IPrimitive) {
	for _, root := range treeList {
		visibleList = append(visibleList, rebuildVisibilityListRecursive(&root, position)...)
	}
	return visibleList
}

func rebuildVisibilityListRecursive(node tree.INode, position mgl32.Vec3) []interfaces.IPrimitive {
	ret := []interfaces.IPrimitive{}
	if node.IsLeaf() == false {
		for _,child := range node.(*tree.Node).Children {
			ret = append(ret, rebuildVisibilityListRecursive(child, position)...)
		}
	} else {
		return node.(*tree.Leaf).FaceList
	}

	return ret
}
