package loader

import (
	entity3 "github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/source-tools-common/entity"
	"github.com/galaco/vmf"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

// Parse Base block.
// Vmf lib is actually capable of doing this;
// contents are loaded into Vmf.Unclassified
func ParseEntities(data string) (vmf.Vmf, error) {
	stringReader := strings.NewReader(data)
	reader := vmf.NewReader(stringReader)

	return reader.Read()
}

func CreateEntity(ent *entity.Entity) entity3.IEntity {
	localEdict := entity3.NewGenericEntity(ent)
	origin := ent.VectorForKey("origin")
	localEdict.Transform().Position = mgl32.Vec3{origin.X(), origin.Y(), origin.Z()}
	localEdict.Transform().Scale = mgl32.Vec3{8, 8, 8}

	return localEdict
}
