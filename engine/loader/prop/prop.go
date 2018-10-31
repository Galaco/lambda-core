package prop

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/material"
	"github.com/galaco/Gource-Engine/engine/mesh"
	"github.com/galaco/Gource-Engine/engine/model"
	studiomodellib "github.com/galaco/Gource-Engine/lib/studiomodel"
	"github.com/galaco/StudioModel"
	"github.com/galaco/StudioModel/mdl"
	"github.com/galaco/StudioModel/phy"
	"github.com/galaco/StudioModel/vtx"
	"github.com/galaco/StudioModel/vvd"
	"strings"
)

func LoadProp(path string) (*model.Model, error) {
	ResourceManager := filesystem.Manager()
	if ResourceManager.Has(path) {
		return ResourceManager.Get(path).(*model.Model), nil
	}
	prop, err := loadProp(strings.Split(path, ".mdl")[0])
	if prop != nil {
		m := modelFromStudioModel(path, prop)
		if m != nil {
			ResourceManager.Add(m)
		}
	} else {
		return ResourceManager.Get(ResourceManager.ErrorModelName()).(*model.Model), err
	}

	return ResourceManager.Get(path).(*model.Model), err
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
	verts, normals, textureCoordinates, err := studiomodellib.VertexDataForModel(studioModel, 0)
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
	for _, dir := range mdlData.TextureDirs {
		for _, name := range mdlData.TextureNames {
			path := strings.Replace(dir, "\\", "/", -1) + name + ".vmt"
			materials = append(materials, material.LoadSingleMaterial(path))
		}
	}
	return materials
}
