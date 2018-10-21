package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/mesh/primitive"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/visibility"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/go-gl/mathgl/mgl32"
)

type World struct {
	entity.Base
	cache       []mesh.IMesh
	faceList    []primitive.IPrimitive
	visData     *visibility.Vis
	LeafCache   *visibility.Cache
	currentLeaf *leaf.Leaf
}

func (entity *World) GetPrimitives() []mesh.IMesh {
	return entity.cache
}

// Rebuild the current facelist to render, by first
// recalculating using vvis data
func (entity *World) UpdateVisibilityList(position mgl32.Vec3) {
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
		entity.cache[0].(*model.Model).Reset()
		entity.cache[0].AddPrimitives(entity.faceList)
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.Cluster {
		return
	}

	entity.LeafCache = entity.visData.GetPVSCacheForCluster(currentLeaf.Cluster)
	if entity.LeafCache != nil {
		primitives := make([]primitive.IPrimitive, 0)
		for _, faceIdx := range entity.LeafCache.Faces {
			primitives = append(primitives, entity.faceList[faceIdx])
		}
		entity.cache[0].(*model.Model).Reset()
		entity.cache[0].AddPrimitives(primitives)
	}
}

func NewWorld(faceList []primitive.IPrimitive, visData *visibility.Vis) *World {
	c := World{
		cache: []mesh.IMesh{
			model.NewModel(make([]primitive.IPrimitive, 0)),
		},
		faceList: faceList,
		visData:  visData,
	}

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
