package loader

import (
	"github.com/galaco/Gource-Engine/core/loader/prop"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/galaco/Gource-Engine/core/resource"
	"github.com/galaco/Gource-Engine/lib/util"
	"github.com/galaco/bsp/primitives/game"
	"strings"
)

// LoadStaticProps GetFile all staticprops referenced in a
// bsp's game lump
func LoadStaticProps(propLump *game.StaticPropLump) []model.StaticProp {
	ResourceManager := resource.Manager()
	prop.LoadProp(ResourceManager.ErrorModelName())

	propPaths := make([]string, 0)
	for _, propEntry := range propLump.PropLumps {
		propPaths = append(propPaths, propLump.DictLump.Name[propEntry.GetPropType()])
	}

	propPaths = util.RemoveDuplicatesFromList(propPaths)
	logger.Notice("Found %d staticprops", len(propPaths))

	numLoaded := 0
	for _, path := range propPaths {
		if !strings.HasSuffix(path, ".mdl") {
			path += ".mdl"
		}
		_, err := prop.LoadProp(path)
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
			staticPropList = append(staticPropList, *model.NewStaticProp(propEntry, &propLump.LeafLump, m))
			continue
		}
		// Model missing, use error model
		m = ResourceManager.GetModel(ResourceManager.ErrorModelName())
		staticPropList = append(staticPropList, *model.NewStaticProp(propEntry, &propLump.LeafLump, m))
	}

	return staticPropList
}
