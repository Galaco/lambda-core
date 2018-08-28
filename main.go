package main

import (
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/systems/renderer"
	"github.com/galaco/go-me-engine/valve/bsp"
	"log"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/mathgl/mgl32"
	bsp2 "github.com/galaco/bsp"
	"github.com/galaco/bsp/primitives/visibility"
)

func main() {
	// Build our engine setup
	Application := engine.NewEngine()
	Application.AddManager(&window.Manager{})
	Application.AddManager(&renderer.Manager{})

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	Application.Initialise()

	// special camera entity
	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	// Load a map!
	LoadMap("data/maps/de_dust2.bsp")

	// Run the engine
	Application.Run()
}

func LoadMap(filename string) {
	// BSP
	bspData := bsp.LoadBsp(filename)
	if bspData.GetHeader().Version < 20 {
		log.Fatal("Unsupported BSP Version. Exiting...")
	}

	// Fetch all BSP face data
	bspPrimitives := bsp.LoadMap(bspData)
	for _,primitive := range bspPrimitives {
		// Ensure created primitive is ready on gpu
		if primitive != nil {
			primitive.GenerateGPUBuffer()
		}
	}

	visData := (*bspData.GetLump(bsp2.LUMP_VISIBILITY).GetContents()).GetData().(*visibility.Vis)

	bspTree := tree.BuildTree(bspData)
	bspComponent := components.NewBspComponent(bspTree, bspPrimitives, visData)
	bspComponent.UpdateVisibilityList(mgl32.Vec3{0, 0, 0})

	worldSpawn := factory.NewEntity(&base.Entity{})
	factory.NewComponent(bspComponent, worldSpawn)
}