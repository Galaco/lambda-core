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
)

func main() {
	runtime.LockOSThread()

	// Build our engine setup
	application := engine.Engine{}
	application.AddManager(&window.Manager{})
	application.AddManager(&renderer.Manager{})

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	application.Initialise()

	// special camera entity
	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	// Create an example entity for rendering
	renderableEnt := factory.NewEntity(&base.Entity{})
	renderableComponent := components.NewRenderableComponent()

	// Load bsp data
	bspData := bsp.LoadBsp("ze_bioshock_v6_3_sample.bsp")
	vertices, faceIndices := bsp.GenerateFacesFromBSP(bspData)
	log.Printf("%d vertices found\n", len(vertices))

	// construct renderable component from bsp primitives
	gpures := renderable.NewGPUResource(vertices)
	for _,f := range faceIndices {
		gpures.AddPrimitive(renderable.NewPrimitive([]float32{}, f))
	}
	renderableComponent.AddRenderableResource(gpures)
	factory.NewComponent(renderableComponent, renderableEnt)

	// Run the engine
	application.Run()
}
