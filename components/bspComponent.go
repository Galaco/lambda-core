package components

import (
	"github.com/galaco/go-me-engine/components/renderable"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/valve/vis"
	"github.com/go-gl/mathgl/mgl32"
	"log"
)

// BspComponent essentially extends a renderable component, as its large number
// of primitives require constant visibility culling; most effectively managed by
// this component itself
type BspComponent struct {
	RenderableComponent

	cache            []interfaces.IGPUMesh
	cachedPosition   mgl32.Vec3
	currentClusterId int16
	faceList         []interfaces.IPrimitive
	visData          *vis.Vis
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

	currentLeaf := component.visData.FindCurrentLeaf(component.cachedPosition)
	if currentLeaf != nil {
		// No need to recalculate face list
		if component.currentClusterId == currentLeaf.ClusterId {
			return
		}
		component.currentClusterId = currentLeaf.ClusterId
		log.Printf("Current Cluster id: %d\n", component.currentClusterId)
	} else {
		if component.currentClusterId == -1 {
			return
		}
		component.currentClusterId = -1
	}

	cache := component.visData.GetCacheLeavesForCluster(component.currentClusterId)
	if cache != nil {
		primitives := make([]interfaces.IPrimitive, 0)
		for _,leaf := range cache.Leaves {
			for _,faceId := range leaf.FaceIndexList {
				primitives = append(primitives, component.faceList[faceId])
			}
		}
		component.cache[0].(*renderable.GPUResourceDynamic).Reset()
		component.cache[0].AddPrimitives(primitives)
	} else {
		component.cache[0].(*renderable.GPUResourceDynamic).Reset()
		component.cache[0].AddPrimitives(component.faceList)
	}
}


func NewBspComponent(faceList []interfaces.IPrimitive, visData *vis.Vis) *BspComponent {
	c := BspComponent{
		cache: []interfaces.IGPUMesh{
			renderable.NewGPUResourceDynamic(make([]interfaces.IPrimitive, 0)),
		},
		faceList: faceList,
		cachedPosition: mgl32.Vec3{65536, 65536, 65536},
		visData: visData,
	}
	c.Etype = T_BspComponent

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
