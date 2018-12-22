package loader

import (
	"github.com/galaco/Gource-Engine/engine/core/logger"
	"github.com/galaco/Gource-Engine/engine/loader/prop"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/resource"
	"github.com/galaco/bsp/primitives/game"
	"log"
	"strings"
)

// LoadStaticProps GetFile all staticprops referenced in a
// bsp's game lump
func LoadStaticProps(propLump *game.StaticPropLump) []model.StaticProp {
	ResourceManager := resource.Manager()
	prop.LoadProp(ResourceManager.ErrorModelName())

	log.Println("Loading static props")
	propPaths := make([]string, 0)
	for _, propEntry := range propLump.PropLumps {
		propPaths = append(propPaths, propLump.DictLump.Name[propEntry.GetPropType()])
	}

	propPaths = buildUniquePropList(propPaths)
	logger.Notice("Found %d staticprops", len(propPaths))

	numLoaded := 0
	for _, path := range propPaths {
		if !strings.HasSuffix(path, ".mdl") {
			path += ".mdl"
		}
		_,err := prop.LoadProp(path)
		if err != nil {
			continue
		}
		numLoaded++
	}

	logger.Notice("Loaded %d props, failed to load %d props", numLoaded, len(propPaths)-numLoaded)

	staticPropList := make([]model.StaticProp, 0)

	for _, propEntry := range propLump.PropLumps {
		modelName := propLump.DictLump.Name[propEntry.GetPropType()]
		m := ResourceManager.GetModel(modelName)
		if m != nil {
			staticPropList = append(staticPropList, *createStaticProp(propEntry, &propLump.LeafLump, m.(*model.Model)))
			continue
		}
		// Model missing, use error model
		m = ResourceManager.GetModel(ResourceManager.ErrorModelName())
		staticPropList = append(staticPropList, *createStaticProp(propEntry, &propLump.LeafLump, m.(*model.Model)))
	}

	return staticPropList
}

// buildUniquePropList Build a list of all different prop files.
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

func createStaticProp(prop game.IStaticPropDataLump, propLeafs *game.StaticPropLeafLump, mod *model.Model) *model.StaticProp {
	return model.NewStaticProp(prop, propLeafs, mod)
}
