package main

import (
	"github.com/galaco/Gource-Engine/components"
	"github.com/galaco/Gource-Engine/engine"
	"github.com/galaco/Gource-Engine/engine/base"
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/event"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/interfaces"
	entity2 "github.com/galaco/Gource-Engine/entity"
	"github.com/galaco/Gource-Engine/message/messages"
	"github.com/galaco/Gource-Engine/message/messagetype"
	"github.com/galaco/Gource-Engine/systems/renderer"
	"github.com/galaco/Gource-Engine/systems/window"
	"github.com/galaco/Gource-Engine/valve/file"
	"github.com/galaco/Gource-Engine/valve/libwrapper/vpk"
	"github.com/galaco/Gource-Engine/valve/loaders/bsp"
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/source-tools-common/entity"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
)

func main() {
	// Build our engine setup
	Application := engine.NewEngine()

	LoadConfig()

	RegisterManagers(Application)

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	Application.Initialise()

	// special camera entity
	cameraEnt := factory.NewEntity(&base.Entity{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	// Load a map!
	LoadMap("data/maps/de_dust2.bsp")

	RegisterShutdownMethod(Application)

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

func LoadConfig() {
	_,err := config.Load()
	if err != nil {
		log.Println(err)
	}
}

func RegisterManagers(app *engine.Engine) {
	app.AddManager(&window.Manager{})
	app.AddManager(&renderer.Manager{})
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

//Implement a way of shutting down the engine
func RegisterShutdownMethod(app *engine.Engine) {
	event.GetEventManager().Listen(messagetype.KeyDown, Closeable{app})
}

