package entity

import (
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/core/loader/prop"
	"github.com/galaco/Lambda-Core/core/resource"
	entity2 "github.com/galaco/Lambda-Core/game/entity"
	"strings"
)

// DoesEntityReferenceStudioModel tests if an entity is
// tied to a model (normally prop_* classnames, but not exclusively)
func DoesEntityReferenceStudioModel(ent entity.IEntity) bool {
	return strings.HasSuffix(ent.KeyValues().ValueForKey("model"), ".mdl")
}

// AssignStudioModelToEntity sets a renderable entity's model
func AssignStudioModelToEntity(entity entity.IEntity) {
	modelName := entity.KeyValues().ValueForKey("model")
	if !resource.Manager().HasModel(modelName) {
		m, _ := prop.LoadProp(modelName, nil)
		entity.(entity2.IProp).SetModel(m)
	} else {
		entity.(entity2.IProp).SetModel(resource.Manager().GetModel(modelName))
	}
}
