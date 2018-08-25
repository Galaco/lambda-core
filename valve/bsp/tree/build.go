package tree

import (
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/model"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/node"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/go-gl/mathgl/mgl32"
)

// Build the bsp node tree
func BuildTree(file *bsp.Bsp) []Node {
	models := (*file.GetLump(bsp.LUMP_MODELS).GetContents()).(lumps.Model).GetData().(*[]model.Model)
	nodes := (*file.GetLump(bsp.LUMP_NODES).GetContents()).(lumps.Node).GetData().(*[]node.Node)
	leafs := (*file.GetLump(bsp.LUMP_LEAFS).GetContents()).(lumps.Leaf).GetData().(*[]leaf.Leaf)
	leafFaces := (*file.GetLump(bsp.LUMP_LEAFFACES).GetContents()).(lumps.LeafFace).GetData().(*[]uint16)

	ret := make([]Node, len(*models))
	for idx,rootModel := range *models {
		rootNode := (*nodes)[rootModel.HeadNode]

		root := Node{
			Id: rootModel.HeadNode,
			Min: rootModel.Mins,
			Max: rootModel.Maxs,
		}

		root = *populateNodeIterable(&root, &rootNode, *nodes, *leafs, *leafFaces)

		ret[idx] = root
	}

	return ret
}

// Recursive load for bsp node/leafs
func populateNodeIterable(node *Node, bspNode *node.Node, bspNodes []node.Node, leafs []leaf.Leaf, leafFaces []uint16) *Node {
	for childNum,childIdx := range bspNode.Children {
		// leaf
		if childIdx < 0 {
			// Child is a leaf
			l := leafs[(-1  -childIdx)]
			faceList := make([]uint16, l.NumLeafFaces)
			for i := uint16(0); i < l.NumLeafFaces; i++ {
				faceList[i] = leafFaces[l.FirstLeafFace + i]
			}
			node.AddChild(childNum, &Leaf{
				Id: -childIdx,
				FaceIndexList: faceList,
			})
		} else {
			// Child is another node
			node.Children[childNum] = &Node{
				Id: childIdx,
				Min: mgl32.Vec3{
					float32(bspNodes[childIdx].Mins[0]),
					float32(bspNodes[childIdx].Mins[1]),
					float32(bspNodes[childIdx].Mins[2]),
				},
				Max: mgl32.Vec3{
					float32(bspNodes[childIdx].Maxs[0]),
					float32(bspNodes[childIdx].Maxs[1]),
					float32(bspNodes[childIdx].Maxs[2]),
				},
			}
			populateNodeIterable(node.Children[childNum].(*Node), &bspNodes[childIdx], bspNodes, leafs, leafFaces)
		}
	}

	return node
}

// Populate BSP tree leafs from loaded face data
func PopulateBspTreeFromFaces(rootNodes []Node, faces []interfaces.IPrimitive) []Node {
	for _,root := range rootNodes {
		populateFaceRecursive(&root, faces)
	}

	return rootNodes
}

// Recursive population of faces in Leaf
func populateFaceRecursive(node INode, faces []interfaces.IPrimitive) {
	if !node.IsLeaf() {
		for idx := range node.(*Node).Children {
			 populateFaceRecursive(node.(*Node).Children[idx], faces)
		}
	} else {
		for _,idx := range node.(*Leaf).FaceIndexList {
			node.(*Leaf).AddFace(faces[idx])
		}
	}
}

// for each model
// 	build node tree from root
//	  get node children
//		if either child is a node
//			repeat get node children