package main

import (
	"github.com/galaco/Gource-Engine/config"
	"github.com/galaco/Gource-Engine/engine"
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messages"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messagetype"
	"github.com/galaco/Gource-Engine/engine/core/logger"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/input/keyboard"
	"github.com/galaco/Gource-Engine/engine/resource"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/game"
	"github.com/galaco/Gource-Engine/lib/gameinfo"
	"github.com/galaco/Gource-Engine/renderer"
	"github.com/galaco/Gource-Engine/window"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	// Load GameInfo.txt
	// GameInfo.txt includes fundamental properties about the game
	// and its resources locations
	cfg, err := config.Load("./")
	if err != nil {
		logger.Fatal(err)
	}
	_, err = gameinfo.LoadConfig(cfg.GameDirectory)
	if err != nil {
		logger.Fatal(err)
	}

	// Register GameInfo.txt referenced resource paths
	// Filesystem module needs to know about all the possible resource
	// locations it can search.
	filesystem.RegisterGameResourcePaths(config.Get().GameDirectory, gameinfo.Get())

	// Explicity define fallbacks for missing resources
	// Defaults are defined, but if HL2 assets are not readable, then
	// the default may not be readable
	resource.Manager().SetErrorModelName("models/props/de_dust/du_antenna_A.mdl")
	resource.Manager().SetErrorTextureName("materials/error.vtf")

	// General engine setup
	Application := engine.NewEngine()
	Application.Initialise()

	Application.AddManager(&window.Manager{})
	Application.AddManager(&renderer.Manager{})

	// Game specific setup
	Game := game.CounterstrikeSource{}
	Game.RegisterEntityClasses()

	// Register behaviour that needs to exist outside of game simulation & control
	RegisterShutdownMethod(Application)

	scene.LoadFromFile(config.Get().GameDirectory + "/maps/de_dust2.bsp")
	//scene.LoadFromFile(config.Get().GameDirectory + "/maps/ze_illya_b3.bsp")

	// Start
	Application.SetSimulationSpeed(10)
	Application.Run()

	resource.Manager().Cleanup()
}

// Closeable Simple struct to control engine shutdown utilising the internal event manager
type Closeable struct {
	target *engine.Engine
}

// ReceiveMessage function will shutdown the engine
func (closer Closeable) ReceiveMessage(message message.IMessage) {
	if message.GetType() == messagetype.KeyDown {
		if message.(*messages.KeyDown).Key == keyboard.KeyEscape {
			// Will shutdown the engine at the end of the current loop
			closer.target.Close()
		}
	}
}

//Implement a way of shutting down the engine
func RegisterShutdownMethod(app *engine.Engine) {
	event.GetEventManager().Listen(messagetype.KeyDown, Closeable{app})
}
