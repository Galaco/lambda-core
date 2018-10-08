package tree

import (
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/galaco/bsp/primitives/node"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/go-gl/mathgl/mgl32"
)

// Build the bsp node tree
func BuildTree(file *bsp.Bsp) []Node {
	models := file.GetLump(bsp.LUMP_MODELS).(*lumps.Model).GetData()
	nodes := file.GetLump(bsp.LUMP_NODES).(*lumps.Node).GetData()
	leafs := file.GetLump(bsp.LUMP_LEAFS).(*lumps.Leaf).GetData()
	leafFaces := file.GetLump(bsp.LUMP_LEAFFACES).(*lumps.LeafFace).GetData()
	planes := file.GetLump(bsp.LUMP_PLANES).(*lumps.Planes).GetData()

	ret := make([]Node, len(models))
	for idx, rootModel := range models {
		rootNode := nodes[rootModel.HeadNode]

		root := Node{
			Id:    rootModel.HeadNode,
			Min:   rootModel.Mins,
			Max:   rootModel.Maxs,
			Plane: &planes[rootNode.PlaneNum],
		}

		root = *populateNodeIterable(&root, &rootNode, nodes, leafs, leafFaces, planes)

		ret[idx] = root
	}

	return ret
}

// Recursive load for bsp node/leafs
func populateNodeIterable(node *Node, bspNode *node.Node, bspNodes []node.Node, leafs []leaf.Leaf, leafFaces []uint16, planes []plane.Plane) *Node {
	for childNum, childIdx := range bspNode.Children {
		// leaf
		if childIdx < 0 {
			// Child is a leaf
			l := leafs[(-1 - childIdx)]
			skyVisible := false
			if l.Flags()&(leaf.LEAF_FLAGS_SKY|leaf.LEAF_FLAGS_SKY2D) != 0 {
				skyVisible = true
			}
			faceIndexList := make([]uint16, l.NumLeafFaces)
			for i := uint16(0); i < l.NumLeafFaces; i++ {
				faceIndexList[i] = leafFaces[l.FirstLeafFace+i]
			}

			node.AddChild(childNum, &Leaf{
				Id:            -1 - childIdx,
				FaceIndexList: faceIndexList,
				ClusterId:     l.Cluster,
				Min: mgl32.Vec3{
					float32(l.Mins[0]),
					float32(l.Mins[1]),
					float32(l.Mins[2]),
				},
				Max: mgl32.Vec3{
					float32(l.Maxs[0]),
					float32(l.Maxs[1]),
					float32(l.Maxs[2]),
				},
				SkyVisible: skyVisible,
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
				Plane: &(planes[bspNodes[childIdx].PlaneNum]),
			}
			populateNodeIterable(node.Children[childNum].(*Node), &bspNodes[childIdx], bspNodes, leafs, leafFaces, planes)
		}
	}

	return node
}
