package bsp

import (
	"github.com/galaco/StudioModel"
	"github.com/galaco/StudioModel/mdl"
	"github.com/galaco/StudioModel/phy"
	"github.com/galaco/StudioModel/vtx"
	"github.com/galaco/StudioModel/vvd"
	"github.com/galaco/bsp/primitives/game"
	"github.com/galaco/go-me-engine/valve/file"
	"log"
	"strings"
)

func LoadStaticProps(propLump *game.StaticPropLump) {
	log.Println("Loading static props")
	propPaths := make([]string, 0)
	for _,propEntry := range propLump.PropLumps {
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
	retList := make([]string, 0)
	for _,entry := range propList {
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
	prop := studiomodel.NewStudioModel(filePath)

	// MDL
	f,err := file.Load(filePath + ".mdl")
	if err != nil {
		log.Println(err)
		return nil
	}
	mdlFile,err := mdl.ReadFromStream(f)
	if err != nil {
		log.Println(err)
		return nil
	}
	prop.AddMdl(mdlFile)


	// VVD
	f,err = file.Load(filePath + ".vvd")
	if err != nil {
		debug.Log(err)
		return nil
	}
	vvdFile, err := vvd.ReadFromStream(f)
	if err != nil {
		debug.Log(err)
		return nil
	}
	prop.AddVvd(vvdFile)

	// VTX
	f,err = file.Load(filePath + ".sw.vtx")
	if err != nil {
		log.Println(err)
		return nil
	}
	vtxFile,err := vtx.ReadFromStream(f)
	if err != nil {
		log.Println(err)
		return nil
	}
	prop.AddVtx(vtxFile)

	// PHY
	f,err = file.Load(filePath + ".phy")
	if err != nil {
		log.Println(err)
	} else {
		phyFile,err := phy.ReadFromStream(f)
		if err != nil {
			log.Println(err)
		}
		prop.AddPhy(phyFile)
	}

	return prop
}
