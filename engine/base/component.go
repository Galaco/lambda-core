package base

import (
	"github.com/galaco/Gource-Engine/engine/core"
)

// A component represents a collection of properties and/or behaviour associated with
// an entity. Common examples would include a renderable component, collision component
// or an AI behaviour component.
// Its common for components behaviour to be controlled by a manager/system, although
// this may not always be necessary.
// THIS component is entirely generic, with empty implementations for the interface
// type. Embed this component into custom components and override methods as necessary.
type Component struct {
	handle core.Handle
	owner  core.Handle
}

// Sets a unique handle for the component.
// Allows for referencing and finding this component easily
func (component *Component) SetHandle(handle core.Handle) {
	component.handle = handle
}

// Runs and setup required by this component
func (component *Component) Initialize() {

}

// Returns this objects unique handle
func (component *Component) GetHandle() core.Handle {
	return component.handle
}

// Returns the handle of the entity this component is
// attached to.
func (component *Component) GetOwnerHandle() core.Handle {
	return component.owner
}

// Set the handle of the owner that this component is
// attached to.
func (component *Component) SetOwnerHandle(handle core.Handle) {
	component.owner = handle
}

// This is called each update loop by the engine
// If this component has any need to control its behaviour, it'll
// most likely be implemented here.
func (component *Component) Update(dt float64) {
}

// Called on component removal.
func (component *Component) Destroy() {
}
