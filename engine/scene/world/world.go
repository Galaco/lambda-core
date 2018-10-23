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
	bspModel     model.Model
	visibleProps  []*StaticProp
	staticProps   []StaticProp
	visData      *visibility.Vis
	LeafCache    *visibility.Cache
	currentLeaf  *leaf.Leaf
}

func (entity *World) GetVisibleBsp() *model.Model {
	return entity.visibleModel
}

func (entity *World) GetVisibleStaticProps() []*StaticProp {
	return entity.visibleProps
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
		if len(entity.visibleModel.GetMeshes()) == len(entity.bspModel.GetMeshes()) {
			return
		}
		entity.visibleModel = &entity.bspModel
		return
	}

	// Haven't changed cluster
	if entity.LeafCache != nil && entity.LeafCache.ClusterId == currentLeaf.Cluster {
		return
	}

	entity.LeafCache = entity.visData.GetPVSCacheForCluster(currentLeaf.Cluster)
	if entity.LeafCache != nil {
		primitives := make([]mesh.IMesh, 0)
		// Rebuild bsp faces
		for _, faceIdx := range entity.LeafCache.Faces {
			primitives = append(primitives, entity.bspModel.GetMeshes()[faceIdx])
		}
		entity.visibleModel = model.NewModel("worldspawn_visible")
		entity.visibleModel.AddMesh(primitives...)

		// Rebuild visible props
		entity.visibleProps = make([]*StaticProp, 0)
		for idx,prop := range entity.staticProps {
			found := false
			for _,leafId := range entity.LeafCache.Leafs {

				entity.visibleProps = append(entity.visibleProps, &entity.staticProps[idx])
				found = true
				break
				
				for _,propLeafId := range prop.leafList {
					if leafId == propLeafId {
						entity.visibleProps = append(entity.visibleProps, &entity.staticProps[idx])
						found = true
						break
					}
				}
				if found == true {
					break
				}
			}
		}
	}
}

func NewWorld(world model.Model, staticProps []StaticProp, visData *visibility.Vis) *World {
	c := World{
		visibleModel: model.NewModel("worldspawn_visible"),
		bspModel: world,
		visibleProps: make([]*StaticProp, 0),
		staticProps: staticProps,
		visData:  visData,
	}

	c.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	return &c
}
