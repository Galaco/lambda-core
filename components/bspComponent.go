package components

import (
	"github.com/galaco/bsp/primitives/visibility"
	"github.com/galaco/go-me-engine/components/bsp"
	"github.com/galaco/go-me-engine/components/renderable"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
	"log"
)

// BspComponent essentially extends a renderable component, as its large number
// of primitives require constant visibility culling; most effectively managed by
// this component itself
type BspComponent struct {
	RenderableComponent
	nodeTrees    []tree.Node
	leafClusters map[int16][]*tree.Leaf
	clusterFaces map[int16][]uint16

	cache          []interfaces.IGPUMesh
	cachedPosition mgl32.Vec3
	currentClusterId int16
	visibilityLump *visibility.Vis
	faceList       []interfaces.IPrimitive
}

// BspCompo
func (component *BspComponent) GetRenderables() []interfaces.IGPUMesh {
	return component.cache
}

// Rebuild the current facelist to render, by first
// recalculating using vvis data
func (component *BspComponent) UpdateVisibilityList(position mgl32.Vec3) {
	// View hasn't moved
	if position.ApproxEqual(component.cachedPosition) {
		return
	}
	component.cachedPosition = position

	currentLeaf := bsp.FindCurrentLeaf(component.nodeTrees, component.cachedPosition)
	if currentLeaf != nil {
		// No need to recalculate face list
		if component.currentClusterId == currentLeaf.ClusterId {
			return
		}
		component.currentClusterId = currentLeaf.ClusterId
		log.Printf("Current Cluster id: %d\n", component.currentClusterId)
	} else {
		component.currentClusterId = -1
	}
	// skip visibility calculation results
	skipVisibilityCulling := true

	// If current == nil then we are outside the map. No visibility calculation.
	// everything is visible
	if component.currentClusterId != -1 && skipVisibilityCulling == false {
		// rebuild facelist for all visible clusters
		visibleLeafFaces := make([]interfaces.IPrimitive, 0)
		for _,clusterId := range component.visibilityLump.GetVisibleClusters(component.currentClusterId) {
			for _,leaf := range component.leafClusters[int16(clusterId)] {
				visibleLeafFaces = append(visibleLeafFaces, leaf.Faces...)
			}
		}

		component.cache[0].(*renderable.GPUResourceDynamic).Reset()
		component.cache[0].AddPrimitives(visibleLeafFaces)
	} else {
		component.cache[0].(*renderable.GPUResourceDynamic).Reset()
		component.cache[0].AddPrimitives(component.faceList)
	}
}

func (component *BspComponent) recursiveBuildClusterList(node tree.INode) {
	if node.IsLeaf() {
		l := node.(*tree.Leaf)
		if component.clusterFaces[l.ClusterId] == nil {
			component.clusterFaces[l.ClusterId] = []uint16{}
		}
		component.clusterFaces[l.ClusterId] = append(component.clusterFaces[l.ClusterId], l.FaceIndexList...)


		if component.leafClusters[l.ClusterId] == nil {
			component.leafClusters[l.ClusterId] = []*tree.Leaf{}
		}
		component.leafClusters[l.ClusterId] = append(component.leafClusters[l.ClusterId], l)
	} else {
		for _, child := range node.(*tree.Node).Children {
			component.recursiveBuildClusterList(child)
		}
	}
}

func NewBspComponent(bspTrees []tree.Node, faceList []interfaces.IPrimitive, visibilityLump *visibility.Vis) *BspComponent {
	c := BspComponent{
		nodeTrees:      bspTrees,
		leafClusters:   map[int16][]*tree.Leaf{},
		clusterFaces: map[int16][]uint16{},
		visibilityLump: visibilityLump,
		cache: []interfaces.IGPUMesh{
			renderable.NewGPUResourceDynamic([]interfaces.IPrimitive{}),
		},
		faceList: faceList,
	}
	c.Etype = T_BspComponent

	for _, root := range c.nodeTrees {
		c.recursiveBuildClusterList(&root)
	}
	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
