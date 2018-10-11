package entity

import (
	"github.com/galaco/Gource-Engine/components/renderable"
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/interfaces"
	"github.com/galaco/Gource-Engine/valve/vis"
	"github.com/go-gl/mathgl/mgl32"
)

type WorldSpawn struct {
	ValveEntity

	cache          []interfaces.IGPUMesh
	cachedPosition mgl32.Vec3
	faceList       []interfaces.IPrimitive
	visData        *vis.Vis
	LeafCache      *vis.Cache
}

func (entity *WorldSpawn) GetPrimitives() []interfaces.IGPUMesh {
	return entity.cache
}

// Rebuild the current facelist to render, by first
// recalculating using vvis data
func (entity *WorldSpawn) UpdateVisibilityList(position mgl32.Vec3) {
	// View hasn't moved
	currentLeaf := entity.visData.FindCurrentLeaf(position)
	if currentLeaf == nil || currentLeaf.ClusterId == -1 {
		// Still outside the world
		if len(entity.cache[0].GetPrimitives()) == len(entity.faceList) {
			return
		}
		entity.cache[0].(*renderable.GPUResourceDynamic).Reset()
		entity.cache[0].AddPrimitives(entity.faceList)
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.ClusterId {
		return
	}
	debug.Logf("Current Cluster id: %d", currentLeaf.ClusterId)

	entity.LeafCache = entity.visData.GetCacheLeavesForCluster(currentLeaf.ClusterId)
	if entity.LeafCache != nil {
		primitives := make([]interfaces.IPrimitive, 0)
		for _, leaf := range entity.LeafCache.Leaves {
			for _, faceIdx := range leaf.FaceIndexList {
				primitives = append(primitives, entity.faceList[faceIdx])
			}
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
		cachedPosition: mgl32.Vec3{65536, 65536, 65536},
		visData:        visData,
	}

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
