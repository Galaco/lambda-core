package prop

import (
	"github.com/galaco/Gource-Engine/core/filesystem"
	material2 "github.com/galaco/Gource-Engine/core/loader/material"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/mesh"
	"github.com/galaco/Gource-Engine/core/model"
	"github.com/galaco/Gource-Engine/core/resource"
	studiomodellib "github.com/galaco/Gource-Engine/lib/studiomodel"
	"github.com/galaco/StudioModel"
	"github.com/galaco/StudioModel/mdl"
	"github.com/galaco/StudioModel/phy"
	"github.com/galaco/StudioModel/vtx"
	"github.com/galaco/StudioModel/vvd"
	"strings"
)

// @TODO This is SUPER incomplete
// Right now it does the bare minimum, and many models seem to have
// some corruption.

// LoadProp loads a single prop/model of known filepath
func LoadProp(path string) (*model.Model, error) {
	ResourceManager := resource.Manager()
	if ResourceManager.HasModel(path) {
		return ResourceManager.GetModel(path).(*model.Model), nil
	}
	prop, err := loadProp(strings.Split(path, ".mdl")[0])
	if prop != nil {
		m := modelFromStudioModel(path, prop)
		if m != nil {
			ResourceManager.AddModel(m)
		} else {
			return ResourceManager.GetModel(ResourceManager.ErrorModelName()).(*model.Model), err
		}
	} else {
		return ResourceManager.GetModel(ResourceManager.ErrorModelName()).(*model.Model), err
	}

	return ResourceManager.GetModel(path).(*model.Model), err
}

func loadProp(filePath string) (*studiomodel.StudioModel, error) {
	prop := studiomodel.NewStudioModel(filePath)

	// MDL
	f, err := filesystem.GetFile(filePath + ".mdl")
	if err != nil {
		return nil, err
	}
	mdlFile, err := mdl.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddMdl(mdlFile)

	// VVD
	f, err = filesystem.GetFile(filePath + ".vvd")
	if err != nil {
		return nil, err
	}
	vvdFile, err := vvd.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddVvd(vvdFile)

	// VTX
	f, err = filesystem.GetFile(filePath + ".dx90.vtx")
	if err != nil {
		return nil, err
	}
	vtxFile, err := vtx.ReadFromStream(f)

	if err != nil {
		return nil, err
	}
	prop.AddVtx(vtxFile)

	// PHY
	f, err = filesystem.GetFile(filePath + ".phy")
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
		logger.Error(err)
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
			path := strings.Replace(dir, "\\", "/", -1) + name + filesystem.ExtensionVmt
			materials = append(materials, material2.LoadSingleMaterial(path))
		}
	}
	return materials
}
