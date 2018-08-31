package engine

import (
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"runtime"
	"time"
)

// Game engine
// Only 1 can be initialised
type engine struct {
	EventManager event.Manager
	Managers     []interfaces.IManager
	Running      bool

	entities   []interfaces.IEntity
	components []interfaces.IComponent
}

// Initialise the engine, and attached managers
func (engine *engine) Initialise() {

	for _, manager := range engine.Managers {
		manager.Register()
	}

}

// Run the engine
func (engine *engine) Run() {
	engine.Running = true

	// Begin the event manager thread in the background
	event.GetEventManager().RunConcurrent()
	// Anything that runs concurrently can start now
	for _, manager := range engine.Managers {
		manager.RunConcurrent()
	}

	dt := 0.0
	startingTime := time.Now().UTC()

	for engine.Running == true {
		for _, manager := range engine.Managers {
			manager.Update(dt)
		}

		for _, manager := range engine.Managers {
			manager.PostUpdate()
		}

		dt = float64(time.Now().UTC().Sub(startingTime).Nanoseconds() / 1000000) / 1000
		startingTime = time.Now().UTC()
	}
}

// Add a new manager to the engine
func (engine *engine) AddManager(manager interfaces.IManager) {
	engine.Managers = append(engine.Managers, manager)
}

func NewEngine() *engine {
	runtime.LockOSThread()
	return &engine{}
}
