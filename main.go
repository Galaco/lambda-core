package main

import (
	"github.com/galaco/Gource-Engine/components"
	"github.com/galaco/Gource-Engine/engine"
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/core/event"
	entity3 "github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/message/messages"
	"github.com/galaco/Gource-Engine/message/messagetype"
	"github.com/galaco/Gource-Engine/systems/renderer"
	"github.com/galaco/Gource-Engine/systems/window"
	"github.com/galaco/Gource-Engine/valve/libwrapper/gameinfo"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	// Build our engine setup
	Application := engine.NewEngine()

	// Load engine configuration
	LoadConfig()
	// Derive and register game resource paths
	gameinfo.RegisterGameResourcePaths(config.Get().GameDirectory, gameinfo.Get())

	RegisterManagers(Application)

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	Application.Initialise()

	// special camera entity - this needs to be refactored out
	cameraEnt := factory.NewEntity(&entity3.Base{})
	factory.NewComponent(components.NewCameraComponent(), cameraEnt)

	// Load a map!
	scene.LoadFromFile(config.Get().GameDirectory + "/maps/de_dust2.bsp")

	// Register behaviour that needs to exist outside of game simulation & control
	RegisterShutdownMethod(Application)

	Application.SetSimulationSpeed(2.5)

	// Run the engine
	Application.Run()
}

// Load project config, then derived game information
func LoadConfig() {
	cfg, err := config.Load()
	if err != nil {
		debug.Log(err)
	}
	gameinfo.LoadConfig(cfg.GameDirectory)
}

func RegisterManagers(app *engine.Engine) {
	app.AddManager(&window.Manager{})
	app.AddManager(&renderer.Manager{})
}

// Simple object to control engine shutdown utilising the internal event manager
type Closeable struct {
	target *engine.Engine
}

func (closer Closeable) ReceiveMessage(message event.IMessage) {
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
