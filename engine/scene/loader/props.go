package loader

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
	"github.com/galaco/Gource-Engine/engine/scene/world"
	studiomodellib "github.com/galaco/Gource-Engine/lib/studiomodel"
	"github.com/galaco/StudioModel"
	"github.com/galaco/StudioModel/mdl"
	"github.com/galaco/StudioModel/phy"
	"github.com/galaco/StudioModel/vtx"
	"github.com/galaco/StudioModel/vvd"
	"github.com/galaco/bsp/primitives/game"
	"log"
	"strings"
)

func LoadStaticProps(propLump *game.StaticPropLump) []world.StaticProp {
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
		if ResourceManager.Has(path) {
			continue
		}
		prop, err := loadProp(strings.Split(path, ".mdl")[0])
		if prop != nil {
			m := modelFromStudioModel(path, prop)
			if m != nil {
				ResourceManager.Add(m)
			}
			numLoaded++
		}
		if err != nil {
			if prop == nil {
				debug.Logf("Failed to load %s, reason was %s", path, err)
			} else {
				debug.Log(err)
			}
			continue
		}
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
		if !ResourceManager.Has("models/error.mdl") {
			prop,err := loadProp("models/error")
			if err != nil{
				continue
			}
			m := modelFromStudioModel("models/error.mdl", prop)
			if m != nil {
				ResourceManager.Add(m)
			}
		}
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

func loadProp(filePath string) (*studiomodel.StudioModel, error) {
	prop := studiomodel.NewStudioModel(filePath)

	// MDL
	f, err := filesystem.Load(filePath + ".mdl")
	if err != nil {
		return nil, err
	}
	mdlFile, err := mdl.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddMdl(mdlFile)

	// VVD
	f, err = filesystem.Load(filePath + ".vvd")
	if err != nil {
		return nil, err
	}
	vvdFile, err := vvd.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddVvd(vvdFile)

	// VTX
	f, err = filesystem.Load(filePath + ".dx90.vtx")
	if err != nil {
		return nil, err
	}
	vtxFile, err := vtx.ReadFromStream(f)

	if err != nil {
		return nil, err
	}
	prop.AddVtx(vtxFile)

	// PHY
	f, err = filesystem.Load(filePath + ".phy")
	if err != nil {
		return prop, err
	}

	phyFile, err := phy.ReadFromStream(f)
	if err != nil {
		return prop, err
	}
	prop.AddPhy(phyFile)

	return prop, nil
}

func modelFromStudioModel(filename string, studioModel *studiomodel.StudioModel) *model.Model {
	verts, normals, textureCoordinates,err := studiomodellib.VertexDataForModel(studioModel, 0)
	if err != nil {
		debug.Log(err)
		return nil
	}
	outModel := model.NewModel(filename)
	mats := materialsForStudioModel(studioModel.Mdl)
	for i := 0; i < len(verts); i++ { //verts is a slice of slices, (ie vertex data per mesh)
		smMesh := mesh.NewMesh()
		smMesh.AddVertex(verts[i]...)
		smMesh.AddNormal(normals[i]...)
		smMesh.AddTextureCoordinate(textureCoordinates[i]...)
		smMesh.Finish()

		//@TODO Map ALL materials to mesh data
		smMesh.SetMaterial(mats[0])

		outModel.AddMesh(smMesh)
	}

	return outModel
}

func materialsForStudioModel(mdlData *mdl.Mdl) []material.IMaterial {
	materials := make([]material.IMaterial, 0)
	for _,dir := range mdlData.TextureDirs {
		for _, name := range mdlData.TextureNames {
			path := strings.Replace(dir, "\\", "/", -1) + name + ".vmt"
			materials = append(materials, material.LoadSingleMaterial(path))
		}
	}
	return materials
}

func createStaticProp(prop game.IStaticPropDataLump, propLeafs *game.StaticPropLeafLump, model *model.Model) *world.StaticProp {
	return world.NewStaticProp(prop, propLeafs, model)
}
