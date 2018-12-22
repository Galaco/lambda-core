package entity

type EntityId uint

var entityHandleCounter EntityId

func newEntityHandle() EntityId {
	entityHandleCounter++
	return entityHandleCounter
}
