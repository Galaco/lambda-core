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

var triangle = []float32{
	0, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
}

var cubeVertices = []float32{
	//  X, Y, Z, U, V
	// Bottom
	-1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, 1.0,
	-1.0, -1.0, 1.0,

	// Top
	-1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,
	-1.0, 1.0, 1.0,
	1.0, 1.0, 1.0,

	// Front
	-1.0, -1.0, 1.0,
	1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, 1.0,
	-1.0, 1.0, 1.0,

	// Back
	-1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, -1.0, -1.0,
	1.0, -1.0, -1.0,
	-1.0, 1.0, -1.0,
	1.0, 1.0, -1.0,

	// Left
	-1.0, -1.0, 1.0,
	-1.0, 1.0, -1.0,
	-1.0, -1.0, -1.0,
	-1.0, -1.0, 1.0,
	-1.0, 1.0, 1.0,
	-1.0, 1.0, -1.0,

	// Right
	1.0, -1.0, 1.0,
	1.0, -1.0, -1.0,
	1.0, 1.0, -1.0,
	1.0, -1.0, 1.0,
	1.0, 1.0, -1.0,
	1.0, 1.0, 1.0,
}

func main() {
	runtime.LockOSThread()
	bspData := bsp.LoadBsp("ze_bioshock_v6_3_sample.bsp")
	log.Println("Loaded bsp.")

	vertexes,indices := bsp.GenerateTrianglesFromBSP(bspData)
	expVerts := make([]float32, len(indices) * 3)

	for i,index := range indices {
		expVerts[(i*3)] = vertexes[index]
		expVerts[(i*3)+1] = vertexes[index+1]
		expVerts[(i*3)+2] = vertexes[index+2]
	}

	log.Printf("%d vertexes found\n", len(vertexes))



	application := engine.Engine{}
	application.AddManager(&window.Manager{})
	application.AddManager(&renderer.Manager{})

	application.Initialise()

	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	renderableEnt := factory.NewEntity(&base.Entity{})
	renderableComponent := components.NewRenderableComponent()
	renderableComponent.SetRenderableResource(renderable.NewGPUResource(expVerts))
	factory.NewComponent(renderableComponent, renderableEnt)


	application.Run()
}
