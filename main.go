package main

import (
	bsp2 "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/systems/renderer"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/valve/bsp"
	"github.com/galaco/go-me-engine/valve/bsp/tree"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
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

	//Implement a way of shutting down the engine
	event.GetEventManager().Listen(messagetype.KeyDown, Closeable{Application})

	Application.SetSimulationSpeed(2.5)

	// Run the engine
	Application.Run()
}

func LoadMap(filename string) {
	// BSP
	bspData := bsp.LoadBsp(filename)
	if bspData.GetHeader().Version < 19 {
		log.Fatal("Unsupported BSP Version. Exiting...")
	}

	// Fetch all BSP face data
	bspPrimitives := bsp.LoadMap(bspData)
	log.Println("Loaded map data")
	for _, primitive := range bspPrimitives {
		// Ensure created primitive is ready on gpu
		if primitive != nil {
			primitive.GenerateGPUBuffer()
		}
	}

	log.Println("Building visibility cluster tree")
	visData := bspData.GetLump(bsp2.LUMP_VISIBILITY).(*lumps.Visibility).GetData()

	worldSpawn := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewBspComponent(tree.BuildTree(bspData), bspPrimitives, visData), worldSpawn)
	log.Println("Cluster tree built.")
}



// Simple object to control engine shutdown utilising the internal event manager
type Closeable struct {
	target *engine.Engine
}

func (closer Closeable) ReceiveMessage(message interfaces.IMessage) {
	if message.GetType() == messagetype.KeyDown {
		if message.(*messages.KeyDown).Key == glfw.KeyEscape {
			// Will shutdown the engine at the end of the current loop
			closer.target.Close()
		}
	}
}
