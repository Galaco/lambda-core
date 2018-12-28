package behaviour

import (
	"github.com/galaco/Gource-Engine/client/input/keyboard"
	"github.com/galaco/Gource-Engine/client/messages"
	"github.com/galaco/Gource-Engine/core"
	"github.com/galaco/Gource-Engine/core/event"
)

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

func NewCloseable(target *core.Engine) *Closeable {
	return &Closeable{
		target: target,
	}
}