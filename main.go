package main

import (
	"github.com/galaco/bsp-viewer/engine"
	"github.com/galaco/bsp-viewer/systems/window"
	"github.com/galaco/bsp-viewer/engine/entity"
	"github.com/galaco/bsp-viewer/components"
	"github.com/galaco/bsp-viewer/engine/factory"
	"github.com/galaco/bsp-viewer/systems/renderer"
)

func main() {
	viewer := engine.Engine{}

	viewer.AddManager(&window.Manager{})
	viewer.AddManager(&renderer.Manager{})

	viewer.Initialise()

	sampleEnt := factory.NewEntity(&entity.Entity{})
	factory.NewComponent(&components.CameraComponent{}, sampleEnt)



	viewer.Run()
}
