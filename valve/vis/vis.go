package vis

import (
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/visibility"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/valve/vis/tree"
	"github.com/go-gl/mathgl/mgl32"
)

type Vis struct {
	ClusterCache []Cache
	BspTree []tree.Node
	VisibilityLump *visibility.Vis
}

func (vis *Vis) GetCacheLeavesForCluster(clusterId int16) *Cache {
	if clusterId == -1 {
		return nil
	}
	for _,cacheEntry := range vis.ClusterCache {
		if cacheEntry.ClusterId == clusterId {
			return &cacheEntry
		}
	}
	return vis.cachePVSForCluster(clusterId)
}

func (vis *Vis) cachePVSForCluster(clusterId int16) *Cache {
	clusterList := vis.VisibilityLump.GetVisibleClusters(clusterId)

	cache := Cache{
		ClusterId: clusterId,
		Leaves: vis.recursiveGetLeavesForPVS(&vis.BspTree[0], clusterList),
	}
	vis.ClusterCache = append(vis.ClusterCache, cache)

	return &cache
}

func (vis *Vis) recursiveGetLeavesForPVS(node tree.INode, clusterList []int16) []*tree.Leaf {
	leaves := make([]*tree.Leaf, 0)
	if node.IsLeaf() {
		for _,clusterId := range clusterList {
			if node.(*tree.Leaf).ClusterId == clusterId {
				leaves = append(leaves, node.(*tree.Leaf))
				return leaves
			}
		}
	} else {
		for _, child := range node.(*tree.Node).Children {
			leaves = append(leaves, vis.recursiveGetLeavesForPVS(child, clusterList)...)
		}
	}
	return leaves
}

func (vis *Vis) FindCurrentLeaf(position mgl32.Vec3) *tree.Leaf {
	// offset [0] is always worldspawn
	// see: https://developer.valvesoftware.com/wiki/Source_BSP_File_Format#Model
	return vis.findCurrentLeafRecursive(&vis.BspTree[0], position)
}

func (vis *Vis) findCurrentLeafRecursive(node tree.INode, position mgl32.Vec3) *tree.Leaf {
	// treat as a npde
	if node.IsLeaf() != true {
		localNode := node.(*tree.Node)

		if isPointInFrontOfPlane(position, localNode.Plane) == true {
			return vis.findCurrentLeafRecursive(localNode.Children[0], position)
		} else {
			return vis.findCurrentLeafRecursive(localNode.Children[1], position)
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





func NewVisFromBSP(file *bsp.Bsp) *Vis {
	return &Vis{
		VisibilityLump: file.GetLump(bsp.LUMP_VISIBILITY).(*lumps.Visibility).GetData(),
		BspTree: tree.BuildTree(file, make([]interfaces.IPrimitive, 0)),
	}
}
