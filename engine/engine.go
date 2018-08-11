package engine

import (
	"github.com/galaco/bsp-viewer/engine/interfaces"
	"github.com/bradhe/stopwatch"
)

const FRAMERATE = 16.6666667

type Engine struct {
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

	dt := 0.0
	timer := stopwatch.Start()

	for engine.Running == true {
		for _, manager := range engine.Managers {
			manager.Update(dt)
		}

		// Restart timer
		dt = 1 / float64(timer.Milliseconds())
		//log.Println(dt)
		//log.Printf("Frametime: %f\n", dt)
		timer.Stop()
		timer = stopwatch.Start()
	}
}

// Add a new manager to the engine
func (engine *Engine) AddManager(manager interfaces.IManager) {
	engine.Managers = append(engine.Managers, manager)
}
