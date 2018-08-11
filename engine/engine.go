package engine

import (
	"github.com/galaco/go-me-engine/engine/interfaces"
	"github.com/bradhe/stopwatch"
	"github.com/galaco/go-me-engine/engine/event"
)

const FRAMERATE = 16.6666667

type Engine struct {
	EventManager event.Manager
	Managers []interfaces.IManager
	Running bool

	entities []interfaces.IEntity
	components []interfaces.IComponent
}

// Initialise the engine, and attached managers
func (engine *Engine) Initialise() {
	for _, manager := range engine.Managers {
		manager.Register()
	}

}

// Run the engine
func (engine *Engine) Run() {
	engine.Running = true

	// Begin the event manager thread in the background
	event.GetEventManager().RunConcurrent()
	// Anything that runs concurrently can start now
	for _, manager := range engine.Managers {
		manager.RunConcurrent()
	}

	dt := 0.0
	timer := stopwatch.Start()

	for engine.Running == true {
		for _, manager := range engine.Managers {
			manager.Update(dt)
		}


		for _, manager := range engine.Managers {
			manager.PostUpdate()
		}

		// Restart timer
		dt = 1 / float64(timer.Milliseconds())
		timer.Stop()
		timer = stopwatch.Start()
	}
}

// Add a new manager to the engine
func (engine *Engine) AddManager(manager interfaces.IManager) {
	engine.Managers = append(engine.Managers, manager)
}
