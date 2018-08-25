package components

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/galaco/go-me-engine/components/bsp"
	"github.com/galaco/go-me-engine/components/renderable"
)

// BspComponent essentially extends a renderable component, as its large number
// of primitives require constant visibility culling; most effectively managed by
// this component itself
type BspComponent struct {
	RenderableComponent
	nodeTrees []tree.Node

	cache []interfaces.IGPUMesh
	cachedPosition mgl32.Vec3
}

// BspCompo
func (component *BspComponent) GetRenderables() []interfaces.IGPUMesh {
	return component.cache
}

func (component *BspComponent) UpdateVisibilityList(position mgl32.Vec3) {
	if position.ApproxEqual(component.cachedPosition) {
		return
	}

	component.cachedPosition = position

	if component.cache == nil || len(component.cache) == 0 {
		component.cache = []interfaces.IGPUMesh{
			renderable.NewGPUResourceDynamic(bsp.RebuildVisibilityList(component.nodeTrees, component.cachedPosition)),
		}
	} else {
		component.cache[0].(*renderable.GPUResourceDynamic).Reset()
		component.cache[0].AddPrimitives(bsp.RebuildVisibilityList(component.nodeTrees, component.cachedPosition))
	}
}


func NewBspComponent(bspTrees []tree.Node) *BspComponent{
	c := BspComponent{
		nodeTrees: bspTrees,
	}
	c.Etype = T_BspComponent

	return &c
}