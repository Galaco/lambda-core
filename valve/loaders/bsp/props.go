package bsp

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/valve/file"
	"github.com/galaco/Gource-Engine/valve/loaders/studiomodel"
	"github.com/galaco/Gource-Engine/valve/loaders/studiomodel/vvd"
	"github.com/galaco/bsp/primitives/game"
	"strings"
)

func LoadStaticProps(propLump *game.StaticPropLump) {
	debug.Log("Loading static props")
	propPaths := []string{}
	for _, propEntry := range propLump.PropLumps {
		propPaths = append(propPaths, propLump.DictLump.Name[propEntry.GetPropType()])
	}

	propPaths = buildUniquePropList(propPaths)
	for _, path := range propPaths {
		loadProp(strings.Split(path, ".mdl")[0])
	}
}

// Build a list of all different prop files.
// Removes duplications
func buildUniquePropList(propList []string) []string {
	retList := []string{}
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

func loadProp(filePath string) *studiomodel.StudioModel {
	// VVD
	f, err := file.Load(filePath + ".vvd")
	if err != nil {
		debug.Log(err)
		return nil
	}
	vvdFile, err := vvd.ReadFromStream(f)
	if err != nil {
		debug.Log(err)
		return nil
	}

	// VTX
	//f,err = file.Load(filePath + ".sw.vtx")
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//vtxFile,err := vtx.ReadFromStream(f)
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}

	return &studiomodel.StudioModel{
		Vvd: vvdFile,
		//		Vtx: vtxFile,
	}
}
