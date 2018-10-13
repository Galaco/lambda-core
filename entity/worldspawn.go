package entity

import (
	"github.com/galaco/Gource-Engine/components/renderable"
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/interfaces"
	"github.com/galaco/Gource-Engine/valve/vis"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/go-gl/mathgl/mgl32"
)

type WorldSpawn struct {
	ValveEntity

	cache          []interfaces.IGPUMesh
	faceList       []interfaces.IPrimitive
	visData        *vis.Vis
	LeafCache      *vis.Cache
	currentLeaf    *leaf.Leaf
}

func (entity *WorldSpawn) GetPrimitives() []interfaces.IGPUMesh {
	return entity.cache
}

// Rebuild the current facelist to render, by first
// recalculating using vvis data
func (entity *WorldSpawn) UpdateVisibilityList(position mgl32.Vec3) {
	// View hasn't moved
	currentLeaf := entity.visData.FindCurrentLeaf(position)

	if currentLeaf == entity.currentLeaf {
		return
	}
	entity.currentLeaf = currentLeaf

	if currentLeaf == nil || currentLeaf.Cluster == -1 {
		// Still outside the world
		if len(entity.cache[0].GetPrimitives()) == len(entity.faceList) {
			return
		}
		entity.cache[0].(*renderable.GPUResourceDynamic).Reset()
		entity.cache[0].AddPrimitives(entity.faceList)
		if currentLeaf != nil {
			debug.Logf("Current Cluster id: %d", currentLeaf.Cluster)
		} else {
			debug.Log("Not in leaf")
		}
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.Cluster {
		return
	}
	debug.Logf("Current Cluster id: %d", currentLeaf.Cluster)

	entity.LeafCache = entity.visData.GetPVSCacheForCluster(currentLeaf.Cluster)
	if entity.LeafCache != nil {
		primitives := make([]interfaces.IPrimitive, 0)
		for _,faceIdx := range entity.LeafCache.Faces {
			primitives = append(primitives, entity.faceList[faceIdx])
		}
		entity.cache[0].(*renderable.GPUResourceDynamic).Reset()
		entity.cache[0].AddPrimitives(primitives)
	}
}

func NewWorld(faceList []interfaces.IPrimitive, visData *vis.Vis) *WorldSpawn {
	c := WorldSpawn{
		cache: []interfaces.IGPUMesh{
			renderable.NewGPUResourceDynamic(make([]interfaces.IPrimitive, 0)),
		},
		faceList:       faceList,
		visData:        visData,
	}

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
