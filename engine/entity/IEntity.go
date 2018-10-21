package entity

import (
	"github.com/galaco/Gource-Engine/engine/core"
	"github.com/galaco/source-tools-common/entity"
)

// Base interface
// All game entities need to implement this
type IEntity interface {
	KeyValues() *entity.Entity
	ClassName() string
	SetHandle(core.Handle)
	GetHandle() core.Handle
	GetComponents() []core.Handle
	AddComponent(handle core.Handle)
}
