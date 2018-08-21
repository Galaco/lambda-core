package tree

import (
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/model"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/node"
	"github.com/galaco/bsp/primitives/leaf"
)

func BuildTree(file *bsp.Bsp) []Node {
	ret := []Node{}
	models := (*file.GetLump(bsp.LUMP_MODELS).GetContents()).(lumps.Model).GetData().(*[]model.Model)
	nodes := (*file.GetLump(bsp.LUMP_NODES).GetContents()).(lumps.Node).GetData().(*[]node.Node)
	leafs := (*file.GetLump(bsp.LUMP_LEAFS).GetContents()).(lumps.Leaf).GetData().(*[]leaf.Leaf)
	leafFaces := (*file.GetLump(bsp.LUMP_LEAFFACES).GetContents()).(lumps.LeafFace).GetData().(*[]uint16)

	for _,root := range *models {
		 rootNode := (*nodes)[root.HeadNode]
		 root := Node{}

		populateNodeIterable(&root, &rootNode, *nodes, *leafs, *leafFaces)

		 ret = append(ret, root)
	}

	return ret
}

func populateNodeIterable(node *Node, bspNode *node.Node, bspNodes[]node.Node, leafs []leaf.Leaf, leafFaces []uint16) {
	for childNum,childIdx := range bspNode.Children {
		// leaf
		if childIdx < 0 {
			// Child is a leaf
			leaf := leafs[(-1  -childIdx)]
			faceList := make([]uint16, leaf.NumLeafFaces)
			for i := uint16(0); i < leaf.NumLeafFaces; i++ {
				faceList[i] = leafFaces[leaf.FirstLeafFace + i]
			}
			node.AddChild(childNum, &Leaf{FaceIndexList: faceList})
		} else {
			// Child is another node
			child := bspNodes[childIdx]
			populateNodeIterable(node, &child, bspNodes, leafs, leafFaces)
		}
	}
}

// for each model
// 	build node tree from root
//	  get node children
//		if either child is a node
//			repeat get node children