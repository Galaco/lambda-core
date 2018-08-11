package core

type Handle uint

var handleCounter = Handle(0)

func NewHandle() Handle {
	handleCounter++
	return handleCounter
}
