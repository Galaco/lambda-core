package world

import (
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/visibility"
	"github.com/galaco/bsp/primitives/leaf"
	"github.com/go-gl/mathgl/mgl32"
	"sync"
)

type World struct {
	entity.Base
	bspModel     model.Model
	staticProps  []StaticProp
	sky          Sky

	visibleWorld *VisibleWorld
	visData      *visibility.Vis
	LeafCache    *visibility.Cache
	currentLeaf  *leaf.Leaf

	rebuildMutex sync.Mutex
}

func (entity *World) VisibleWorld() *VisibleWorld {
	entity.rebuildMutex.Lock()
	vw := entity.visibleWorld
	entity.rebuildMutex.Unlock()
	return vw
}

// Rebuild the current facelist to render, by first
// recalculating using vvis data
func (entity *World) TestVisibility(position mgl32.Vec3) {
	// View hasn't moved
	currentLeaf := entity.visData.FindCurrentLeaf(position)

	if currentLeaf == entity.currentLeaf {
		return
	}
	entity.currentLeaf = currentLeaf

	if currentLeaf == nil || currentLeaf.Cluster == -1 {
		// Still outside the world
		if len(entity.visibleWorld.visibleModel.GetMeshes()) == len(entity.bspModel.GetMeshes()) {
			return
		}
		entity.AsyncRebuildVisibleWorld()
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.Cluster {
		return
	}

	entity.LeafCache = entity.visData.GetPVSCacheForCluster(entity.currentLeaf.Cluster)

	entity.AsyncRebuildVisibleWorld()
}

// Launches rebuilding the visible world in a separate thread
// Note: This *could* cause rendering issues if the rebuild is slower than
// travelling between clusters
func (entity *World) AsyncRebuildVisibleWorld() {
	go func(cache *visibility.Cache) {
		visibleWorld := &VisibleWorld{
			sky: &entity.sky,
		}
		if entity.LeafCache != nil {
			primitives := make([]mesh.IMesh, 0)
			// Rebuild bsp faces
			for _, faceIdx := range cache.Faces {
				primitives = append(primitives, entity.bspModel.GetMeshes()[faceIdx])
			}
			visibleModel := model.NewModel("worldspawn_visible")
			visibleModel.AddMesh(primitives...)

			// Rebuild visible props
			visibleProps := make([]*StaticProp, 0)
			for idx, prop := range entity.staticProps {
				found := false
				for _, leafId := range cache.Leafs {
					for _, propLeafId := range prop.leafList {
						if leafId == propLeafId {
							visibleProps = append(visibleProps, &entity.staticProps[idx])
							found = true
							break
						}
					}
					if found == true {
						break
					}
				}
			}

			visibleWorld.visibleModel = visibleModel
			visibleWorld.visibleProps = visibleProps
		} else {
			entity.rebuildMutex.Lock()
			visibleWorld.visibleModel = &entity.bspModel
			entity.rebuildMutex.Unlock()
		}

		entity.rebuildMutex.Lock()
		entity.visibleWorld = visibleWorld
		entity.rebuildMutex.Unlock()
	}(entity.LeafCache)
}

// Build skybox from tree
func (entity *World) BuildSkybox(sky *model.Model, position mgl32.Vec3, scale float32) {
	l := entity.visData.FindCurrentLeaf(position)
	cache := entity.visData.GetPVSCacheForCluster(l.Cluster)

	primitives := make([]mesh.IMesh, 0)
	// Rebuild bsp faces
	for _, faceIdx := range cache.Faces {
		primitives = append(primitives, entity.bspModel.GetMeshes()[faceIdx])
	}
	visibleModel := model.NewModel("worldspawn_visible")
	visibleModel.AddMesh(primitives...)

	// Rebuild visible props
	visibleProps := make([]*StaticProp, 0)
	for idx, prop := range entity.staticProps {
		found := false
		for _, leafId := range cache.Leafs {
			for _, propLeafId := range prop.leafList {
				if leafId == propLeafId {
					visibleProps = append(visibleProps, &entity.staticProps[idx])
					found = true
					break
				}
			}
			if found == true {
				break
			}
		}
	}

	entity.sky = *NewSky(visibleModel, sky, visibleProps, position, scale)
}

func NewWorld(world model.Model, staticProps []StaticProp, visData *visibility.Vis) *World {
	c := World{
		bspModel:     world,
		staticProps:  staticProps,
		visData:      visData,
		visibleWorld: NewVisibleWorld(),
	}

	c.TestVisibility(mgl32.Vec3{0, 0, 0})

	return &c
}
