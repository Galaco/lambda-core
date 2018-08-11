package main

import (
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/systems/renderer"
	"github.com/galaco/go-me-engine/components/renderable"
)

var triangle = []float32{
	0, 0.5, 0,
	-0.5, -0.5, 0,
	0.5, -0.5, 0,
}

func main() {
	viewer := engine.Engine{}

	viewer.AddManager(&window.Manager{})
	viewer.AddManager(&renderer.Manager{})

	viewer.Initialise()

	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	renderableEnt := factory.NewEntity(&base.Entity{})
	renderableComponent := components.NewRenderableComponent()
	renderableComponent.SetRenderableResource(renderable.NewGPUResource(triangle))
	factory.NewComponent(renderableComponent, renderableEnt)


	viewer.Run()
}
