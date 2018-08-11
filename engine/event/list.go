package event

type Id string

func (id *Id) String() string {
	return string(*id)
}

const KeyDown = Id("KeyDown")
const KeyHeld = Id("KeyHeld")
const KeyReleased = Id("KeyReleased")