package bsp

import (
	"github.com/galaco/Gource/components"
	"github.com/galaco/Gource/components/renderable"
	"github.com/galaco/Gource/engine/base/primitive"
	"github.com/galaco/Gource/engine/factory"
	"github.com/galaco/Gource/engine/interfaces"
	entity2 "github.com/galaco/Gource/entity"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/vmf"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

// Parse Entity block.
// Vmf lib is actually capable of doing this;
// contents are loaded into Vmf.Unclassified
func ParseEntities(data string) (vmf.Vmf, error) {
	stringReader := strings.NewReader(data)
	reader := vmf.NewReader(stringReader)

	return reader.Read()
}

func CreateEntity(ent *entity.Entity) {
	localEdict := &entity2.ValveEntity{}
	origin := ent.VectorForKey("origin")
	localEdict.GetTransformComponent().Position = mgl32.Vec3{origin.X(), origin.Y(), origin.Z()}
	localEdict.GetTransformComponent().Scale = mgl32.Vec3{8, 8, 8}

	placeholder := components.NewRenderableComponent()
	resource := renderable.NewGPUResource([]interfaces.IPrimitive{primitive.NewCube()})
	resource.Prepare()
	placeholder.AddRenderableResource(resource)
	factory.NewComponent(placeholder, factory.NewEntity(localEdict))
}
