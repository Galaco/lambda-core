package prop

import (
	"github.com/galaco/Lambda-Core/core/filesystem"
	material2 "github.com/galaco/Lambda-Core/core/loader/material"
	"github.com/galaco/Lambda-Core/core/logger"
	"github.com/galaco/Lambda-Core/core/material"
	"github.com/galaco/Lambda-Core/core/mesh"
	"github.com/galaco/Lambda-Core/core/model"
	"github.com/galaco/Lambda-Core/core/resource"
	studiomodellib "github.com/galaco/Lambda-Core/lib/studiomodel"
	"github.com/galaco/StudioModel"
	"github.com/galaco/StudioModel/mdl"
	"github.com/galaco/StudioModel/phy"
	"github.com/galaco/StudioModel/vtx"
	"github.com/galaco/StudioModel/vvd"
	"strings"
)

// @TODO This is SUPER incomplete
// right now it does the bare minimum, and many models seem to have
// some corruption.

// LoadProp loads a single prop/model of known filepath
func LoadProp(path string, fs filesystem.IFileSystem) (*model.Model, error) {
	ResourceManager := resource.Manager()
	if ResourceManager.HasModel(path) {
		return ResourceManager.Model(path), nil
	}
	prop, err := loadProp(strings.Split(path, ".mdl")[0], fs)
	if prop != nil {
		m := modelFromStudioModel(path, prop, fs)
		if m != nil {
			ResourceManager.AddModel(m)
		} else {
			return ResourceManager.Model(ResourceManager.ErrorModelName()), err
		}
	} else {
		return ResourceManager.Model(ResourceManager.ErrorModelName()), err
	}

	return ResourceManager.Model(path), err
}

func loadProp(filePath string, fs filesystem.IFileSystem) (*studiomodel.StudioModel, error) {
	prop := studiomodel.NewStudioModel(filePath)

	// MDL
	f, err := fs.GetFile(filePath + ".mdl")
	if err != nil {
		return nil, err
	}
	mdlFile, err := mdl.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddMdl(mdlFile)

	// VVD
	f, err = fs.GetFile(filePath + ".vvd")
	if err != nil {
		return nil, err
	}
	vvdFile, err := vvd.ReadFromStream(f)
	if err != nil {
		return nil, err
	}
	prop.AddVvd(vvdFile)

	// VTX
	f, err = fs.GetFile(filePath + ".dx90.vtx")
	if err != nil {
		return nil, err
	}
	vtxFile, err := vtx.ReadFromStream(f)

	if err != nil {
		return nil, err
	}
	prop.AddVtx(vtxFile)

	// PHY
	f, err = fs.GetFile(filePath + ".phy")
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

func modelFromStudioModel(filename string, studioModel *studiomodel.StudioModel, fs filesystem.IFileSystem) *model.Model {
	verts, normals, textureCoordinates, err := studiomodellib.VertexDataForModel(studioModel, 0)
	if err != nil {
		logger.Error(err)
		return nil
	}
	outModel := model.NewModel(filename)
	mats := materialsForStudioModel(studioModel.Mdl, fs)
	for i := 0; i < len(verts); i++ { //verts is a slice of slices, (ie vertex data per mesh)
		smMesh := mesh.NewMesh()
		smMesh.AddVertex(verts[i]...)
		smMesh.AddNormal(normals[i]...)
		smMesh.AddUV(textureCoordinates[i]...)
		//smMesh.Finish()

		//@TODO Map ALL materials to mesh data
		smMesh.SetMaterial(mats[0])

		outModel.AddMesh(smMesh)
	}

	return outModel
}

func materialsForStudioModel(mdlData *mdl.Mdl, fs filesystem.IFileSystem) []material.IMaterial {
	materials := make([]material.IMaterial, 0)
	for _, dir := range mdlData.TextureDirs {
		for _, name := range mdlData.TextureNames {
			path := strings.Replace(dir, "\\", "/", -1) + name + filesystem.ExtensionVmt
			materials = append(materials, material2.LoadSingleMaterial(path, fs))
		}
	}
	return materials
}
