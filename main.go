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
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/source-tools-common/texdatastringtable"
	"github.com/galaco/vtf"
	"github.com/galaco/go-me-engine/components/renderable/material"
	"github.com/galaco/go-me-engine/engine/filesystem"
	vpk2 "github.com/galaco/vpk2"
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

	// Create an example entity for rendering
	renderableEnt := factory.NewEntity(&base.Entity{})
	renderableComponent := components.NewRenderableComponent()


	// Load bsp data
	bspData := bsp.LoadBsp("data/maps/de_dust2.bsp")
	vertices, faceIndices, texInfos := bsp.GenerateFacesFromBSP(bspData)
	log.Printf("%d vertices found\n", len(vertices))

	stringDataLump := *bspData.GetLump(bsplib.LUMP_TEXDATA_STRING_DATA).GetContents()
	stringTableLump := *bspData.GetLump(bsplib.LUMP_TEXDATA_STRING_TABLE).GetContents()
	stringtable := texdatastringtable.NewTable(
		*stringDataLump.GetData().(*string),
		*stringTableLump.(lumps.TexDataStringTable).GetData().(*[]int32))


	// Derive a unique list of all materials referenced in the map
	materialList := []string{}
	for _,ti := range texInfos {
		target,_ := stringtable.GetString(int(ti.TexData))
		found := false
		for _,cur := range materialList {
			if cur == target {
				found = true
				break
			}
		}
		if found == false {
			materialList = append(materialList, target)
		}
	}

	// Load all reference materials into memory
	vpkHandle,err := vpk2.Open(vpk2.MultiVPK("data/cstrike/cstrike_pak"))

	if err != nil {
		log.Fatal(err)
	}
	for _,materialPath := range materialList {
		// Load file from vpk into memory
		vpkFile := vpkHandle.Entry("materials/" + materialPath + ".vtf")
		if vpkFile == nil {
			log.Println("Couldnt find material: materials/" + materialPath + ".vtf")
			continue
		}
		file,err := vpkFile.Open()

		// Lets not fatal error if a texture is missing
		if err != nil {
			continue
		}
		// Convert from vtf to raw rgb
		//s,_ := file.Stat()
		//buf := make([]byte, s.Size())
		//buf.
		texture,err := vtf.ReadFromStream(file)
		// Again if texture is broke, still continue
		if err != nil {
			continue
		}
		//a := texture.GetHighestResolutionImageForFrame(0)
		//log.Println(a)
		// Store file containing raw data in memory
		FileManager.AddFile(
			material.NewMaterial(
				materialPath,
				texture.GetLowResImageData(),
				//texture.GetHighestResolutionImageForFrame(0),
				int(texture.GetHeader().Width),
				int(texture.GetHeader().Height)))
		// Finally generate the gpu buffer for the material
		FileManager.GetFile(materialPath).(*material.Material).GenerateGPUBuffer()
	}

	// construct renderable component from bsp primitives
	gpures := renderable.NewGPUResource(vertices)
	for idx,f := range faceIndices {
		// This is basically creating a primitive for each face, with material
		target,_ := stringtable.GetString(int(texInfos[idx].TexData))
		primitive := renderable.NewPrimitive([]float32{}, f)
		// @TODO Ensure a default material is set when not found
		if FileManager.GetFile(target) != nil {
			uvs := bsp.TexCoordsForFaceFromTexInfo(vertices, &texInfos[idx])
			primitive.AddTextureCoordinateData(uvs)
			primitive.AddMaterial(FileManager.GetFile(target).(*material.Material))
		}
		gpures.AddPrimitive(primitive)
	}
	renderableComponent.AddRenderableResource(gpures)
	factory.NewComponent(renderableComponent, renderableEnt)

	// Run the engine
	Application.Run()
}