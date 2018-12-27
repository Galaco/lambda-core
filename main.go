package main

import (
	"github.com/galaco/Gource-Engine/client/config"
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/client/messages"
	"github.com/galaco/Gource-Engine/client/renderer"
	"github.com/galaco/Gource-Engine/client/scene"
	"github.com/galaco/Gource-Engine/client/window"
	"github.com/galaco/Gource-Engine/core"
	"github.com/galaco/Gource-Engine/core/event"
	"github.com/galaco/Gource-Engine/core/filesystem"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/resource"
	"github.com/galaco/Gource-Engine/game"
	"github.com/galaco/Gource-Engine/lib/gameinfo"
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
	Application := core.NewEngine()
	Application.Initialise()

	windowName := "Gource"
	gameInfoNode, _ := gameinfo.Get().Find("GameInfo")
	if gameInfoNode == nil {
		logger.Fatal("gameinfo was not found.")
	}
	gameNode, _ := gameInfoNode.Find("game")
	if gameNode != nil {
		windowName, _ = gameNode.AsString()
	}
	Application.AddManager(&window.Manager{
		Name: windowName,
	})
	Application.AddManager(&renderer.Manager{})

	// Game specific setup
	Game := game.CounterstrikeSource{}
	Game.RegisterEntityClasses()

	// Register behaviour that needs to exist outside of game simulation & control
	RegisterShutdownMethod(Application)

	scene.LoadFromFile(config.Get().GameDirectory + "/maps/de_dust2.bsp")

	// Start
	Application.SetSimulationSpeed(10)
	Application.Run()

	resource.Manager().Cleanup()
}

// Closeable Simple struct to control engine shutdown utilising the internal event manager
type Closeable struct {
	target *core.Engine
}

// ReceiveMessage function will shutdown the engine
func (closer Closeable) ReceiveMessage(message event.IMessage) {
	if message.GetType() == messages.TypeKeyDown {
		if message.(*messages.KeyDown).Key == keyboard.KeyEscape {
			// Will shutdown the engine at the end of the current loop
			closer.target.Close()
		}
	}
}

//Implement a way of shutting down the engine
func RegisterShutdownMethod(app *core.Engine) {
	event.GetEventManager().Listen(messages.TypeKeyDown, Closeable{app})
}
