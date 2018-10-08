package vis

import (
	"github.com/galaco/Gource/valve/vis/tree"
	"github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/bsp/primitives/plane"
	"github.com/galaco/bsp/primitives/visibility"
	"github.com/go-gl/mathgl/mgl32"
)

type Vis struct {
	ClusterCache   []Cache
	BspTree        []tree.Node
	VisibilityLump *visibility.Vis

	viewPosition    mgl32.Vec3
	viewCurrentLeaf *tree.Leaf
}

func (vis *Vis) GetCacheLeavesForCluster(clusterId int16) *Cache {
	if clusterId == -1 {
		return nil
	}
	for _, cacheEntry := range vis.ClusterCache {
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
		Leaves:    vis.recursiveGetLeavesForPVS(&vis.BspTree[0], clusterList),
	}
	for _, l := range cache.Leaves {
		if l.SkyVisible == true {
			cache.SkyVisible = true
			break
		}
	}
	vis.ClusterCache = append(vis.ClusterCache, cache)

	return &cache
}

func (vis *Vis) recursiveGetLeavesForPVS(node tree.INode, clusterList []int16) []*tree.Leaf {
	leaves := make([]*tree.Leaf, 0)
	if node.IsLeaf() {
		for _, clusterId := range clusterList {
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
	if vis.viewPosition.ApproxEqualThreshold(position, 0.000000001) == false {
		vis.viewPosition = position
		// offset [0] is always worldspawn
		// see: https://developer.valvesoftware.com/wiki/Source_BSP_File_Format#Model
		vis.viewCurrentLeaf = vis.findCurrentLeafRecursive(&vis.BspTree[0], vis.viewPosition)
	}
	return vis.viewCurrentLeaf
}

func (vis *Vis) findCurrentLeafRecursive(node tree.INode, position mgl32.Vec3) *tree.Leaf {
	// treat as a npde
	if node.IsLeaf() != true {
		if isPointInFrontOfPlane(position, node.(*tree.Node).Plane) == true {
			return vis.findCurrentLeafRecursive(node.(*tree.Node).Children[0], position)
		} else {
			return vis.findCurrentLeafRecursive(node.(*tree.Node).Children[1], position)
		}
	}

	return node.(*tree.Leaf)
}

// Check if viewpoint is in front or behind the split plane
// dot product of place to origin & plane to viewpoint
func isPointInFrontOfPlane(point mgl32.Vec3, nodePlane *plane.Plane) bool {
	return point.Dot(nodePlane.Normal) > nodePlane.Distance
}

func NewVisFromBSP(file *bsp.Bsp) *Vis {
	return &Vis{
		VisibilityLump: file.GetLump(bsp.LUMP_VISIBILITY).(*lumps.Visibility).GetData(),
		BspTree:        tree.BuildTree(file),
		viewPosition:   mgl32.Vec3{65536, 65536, 65536},
	}
}
