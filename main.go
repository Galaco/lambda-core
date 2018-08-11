package main

import (
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/engine/entity"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/systems/renderer"
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
