package main

import (
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/go-me-engine/components"
	"github.com/galaco/go-me-engine/engine"
	"github.com/galaco/go-me-engine/engine/base"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/factory"
	"github.com/galaco/go-me-engine/engine/interfaces"
	entity2 "github.com/galaco/go-me-engine/entity"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/galaco/go-me-engine/systems/renderer"
	"github.com/galaco/go-me-engine/systems/window"
	"github.com/galaco/go-me-engine/valve/file"
	"github.com/galaco/go-me-engine/valve/libwrapper/vpk"
	"github.com/galaco/go-me-engine/valve/loaders/bsp"
	"github.com/galaco/source-tools-common/entity"
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
	bspData, err := bsplib.ReadFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if bspData.GetHeader().Version < 20 {
		log.Fatal("Unsupported BSP Version. Exiting...")
	}

	// Setup all possible resource loading locations
	vpkHandle, err := vpk.OpenVPK("data/cstrike/cstrike_pak")
	if err != nil {
		log.Fatal(err)
	}
	file.SetGameVPK(vpkHandle)
	file.SetPakfile(bspData.GetLump(bsplib.LUMP_PAKFILE).(*lumps.Pakfile))

	// Load worldspawn
	worldSpawn := factory.NewEntity(bsp.LoadMap(bspData)).(*entity2.WorldSpawn)


	// Get entdata
	vmfEntityTree, err := bsp.ParseEntities(bspData.GetLump(bsplib.LUMP_ENTITIES).(*lumps.EntData).GetData())
	if err != nil {
		log.Fatal(err)
	}
	entityList := entity.FromVmfNodeTree(vmfEntityTree.Unclassified)
	log.Printf("Found %d entities\n", entityList.Length())
	for i := 0; i < entityList.Length(); i++ {
		bsp.CreateEntity(entityList.Get(i))
	}

	worldSpawn.Definition = entityList.FindByKeyValue("classname", "worldspawn")
	sky, err := bsp.LoadSky(worldSpawn.Definition.ValueForKey("skyname"))
	if err == nil {
		factory.NewComponent(sky, worldSpawn)
	}
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
