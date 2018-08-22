package main

import (
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/systems/renderer"
	"github.com/galaco/go-me-engine/components/renderable"
	"github.com/galaco/go-me-engine/valve/bsp"
	"log"
	"runtime"
	"github.com/galaco/vtf"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	"github.com/galaco/go-me-engine/valve/stringtable"
	"github.com/galaco/go-me-engine/valve/vpk"
	vpk2 "github.com/galaco/vpk2"
	bsp2 "github.com/galaco/bsp"
)

func main() {
	runtime.LockOSThread()

	FileManager := filesystem.GetFileManager()

	// Build our engine setup
	Application := engine.Engine{}
	Application.AddManager(&window.Manager{})
	Application.AddManager(&renderer.Manager{})

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	Application.Initialise()

	// special camera entity
	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	// BSP
	bspData := LoadBSP("data/maps/de_dust2.bsp")
	faceVertices, faceIndices, texInfos, faceNormals := bsp.GenerateFacesFromBSP(bspData)

	// Open VPK filesystem
	vpkHandle,err := vpk.OpenVPK("data/cstrike/cstrike_pak")
	if err != nil {
		log.Fatal(err)
	}

	// MATERIALS
	stringTable := stringtable.GetTable(bspData)
	// Derive a unique list of all materials referenced in the map
	materialList := stringtable.SortUnique(stringTable, texInfos)
	LoadMaterials(vpkHandle, materialList)

	// construct renderable component from bsp primitives
	bspPrimitives := make([]renderable.IPrimitive, len(faceIndices))
	for idx,f := range faceIndices {
		// This is basically creating a primitive for each face, with material
		target,_ := stringTable.GetString(int(texInfos[idx].TexData))
		primitive := renderable.NewPrimitive(faceVertices[idx], f, faceNormals[idx])
		// @TODO Ensure a default material is set when not found
		if FileManager.GetFile(target) != nil {
			mat := FileManager.GetFile(target).(*material.Material)
			primitive.AddMaterial(mat)
			primitive.AddTextureCoordinateData(bsp.TexCoordsForFaceFromTexInfo(faceVertices[idx], &texInfos[idx], mat.GetWidth(), mat.GetHeight()))
		} else {
			primitive.AddTextureCoordinateData(bsp.TexCoordsForFaceFromTexInfo(faceVertices[idx], &texInfos[idx], 1, 1))
		}
		bspPrimitives[idx] = primitive
	}

	// Prepare a renderable component from our bsp primitives
	renderableComponent := components.NewRenderableComponent()
	renderableComponent.AddRenderableResource(renderable.NewGPUResource(bspPrimitives))

	// Add component to an entity
	renderableEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(renderableComponent, renderableEnt)

	// Run the engine
	Application.Run()
}

func LoadBSP(filename string) *bsp2.Bsp {
	f := bsp.LoadBsp(filename)
	if f.GetHeader().Version < 20 {
		log.Fatal("Unsupported BSP Version. Exiting...")
	}

	return f
}

func LoadMaterials(vpkHandle *vpk2.VPK, materialList []string) {
	FileManager := filesystem.GetFileManager()

	for _,materialPath := range materialList {
		// Load file from vpk into memory
		vpkFile := vpkHandle.Entry("materials/" + materialPath + ".vtf")
		if vpkFile == nil {
			log.Println("Couldnt find material: materials/" + materialPath + ".vtf")
			continue
		}
		file,err := vpkFile.Open()

		// Its quite possible for a texture to be missing, just skip it.
		if err != nil {
			continue
		}

		// Attempt to parse the vtf into color data we can use,
		// if this fails (it shouldn't) we can treat it like it was missing
		texture,err := vtf.ReadFromStream(file)
		if err != nil {
			log.Println(err)
			continue
		}
		// Store file containing raw data in memory
		FileManager.AddFile(
			material.NewMaterial(
				materialPath,
				texture,
				int(texture.GetHeader().Width),
				int(texture.GetHeader().Height)))
		// Finally generate the gpu buffer for the material
		FileManager.GetFile(materialPath).(*material.Material).GenerateGPUBuffer()
	}
}