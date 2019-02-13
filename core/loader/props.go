package loader

import (
	"github.com/galaco/Lambda-Core/core/filesystem"
	"github.com/galaco/Lambda-Core/core/loader/prop"
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda-Core/core/model"
	"github.com/galaco/Lambda-Core/core/resource"
	"github.com/galaco/Lambda-Core/lib/util"
	"github.com/galaco/bsp/primitives/game"
	"strings"
)

// LoadStaticProps GetFile all staticprops referenced in a
// bsp's game lump
func LoadStaticProps(propLump *game.StaticPropLump, fs *filesystem.FileSystem) []model.StaticProp {
	ResourceManager := resource.Manager()
	_,err := prop.LoadProp(ResourceManager.ErrorModelName(), fs)
	// If we have no error model, expect this to be fatal issue
	if err != nil {
		logger.Fatal(err)
	}

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
		_, err := prop.LoadProp(path, fs)
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
