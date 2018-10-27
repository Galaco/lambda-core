package loader

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/loader/prop"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	"github.com/galaco/bsp/primitives/game"
	"log"
	"strings"
)

func LoadStaticProps(propLump *game.StaticPropLump) []world.StaticProp {
	prop.LoadProp("models/error.mdl")

	ResourceManager := filesystem.Manager()
	log.Println("Loading static props")
	propPaths := make([]string, 0)
	for _, propEntry := range propLump.PropLumps {
		propPaths = append(propPaths, propLump.DictLump.Name[propEntry.GetPropType()])
	}

	propPaths = buildUniquePropList(propPaths)
	debug.Logf("Found %d staticprops", len(propPaths))

	numLoaded := 0
	for _, path := range propPaths {
		if !strings.HasSuffix(path, ".mdl") {
			path += ".mdl"
		}
		prop.LoadProp(path)
	}

	debug.Logf("Loaded %d props, failed to load %d props", numLoaded, len(propPaths)-numLoaded)

	staticPropList := make([]world.StaticProp, 0)

	for _, propEntry := range propLump.PropLumps {
		modelName := propLump.DictLump.Name[propEntry.GetPropType()]
		m := ResourceManager.Get(modelName)
		if m != nil {
			staticPropList = append(staticPropList, *createStaticProp(propEntry, &propLump.LeafLump, m.(*model.Model)))
			continue
		}
		// Model missing, use error model
		m = ResourceManager.Get("models/error.mdl")
		staticPropList = append(staticPropList, *createStaticProp(propEntry, &propLump.LeafLump, m.(*model.Model)))
	}

	return staticPropList
}

// Build a list of all different prop files.
// Removes duplications
func buildUniquePropList(propList []string) []string {
	retList := make([]string, 0)

	for _, entry := range propList {
		found := false
		for _, unique := range retList {
			if entry == unique {
				found = true
				break
			}
		}
		if !found {
			retList = append(retList, entry)
		}
	}

	return retList
}

func createStaticProp(prop game.IStaticPropDataLump, propLeafs *game.StaticPropLeafLump, model *model.Model) *world.StaticProp {
	return world.NewStaticProp(prop, propLeafs, model)
}
