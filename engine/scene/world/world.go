package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/visibility"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/go-gl/mathgl/mgl32"
)

type World struct {
	entity.Base
	visibleModel *model.Model
	model 	     model.Model
	visData     *visibility.Vis
	LeafCache   *visibility.Cache
	currentLeaf *leaf.Leaf
}

func (entity *World) GetModel() *model.Model {
	return entity.visibleModel
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
		if len(entity.visibleModel.GetMeshes()) == len(entity.model.GetMeshes()) {
			return
		}
		entity.visibleModel = &entity.model
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.Cluster {
		return
	}

	entity.LeafCache = entity.visData.GetPVSCacheForCluster(currentLeaf.Cluster)
	if entity.LeafCache != nil {
		primitives := make([]mesh.IMesh, 0)
		for _, faceIdx := range entity.LeafCache.Faces {
			primitives = append(primitives, entity.model.GetMeshes()[faceIdx])
		}
		entity.visibleModel = model.NewModel()
		entity.visibleModel.AddMesh(primitives...)
	}
}

func NewWorld(world model.Model, visData *visibility.Vis) *World {
	c := World{
		visibleModel: model.NewModel(),
		model: world,
		visData:  visData,
	}

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
