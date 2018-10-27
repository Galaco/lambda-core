package entity

import (
	"github.com/galaco/source-tools-common/entity"
)

// Base interface
// All game entities need to implement this
type IEntity interface {
	KeyValues() *entity.Entity
	SetKeyValues(entity2 *entity.Entity)
	Classname() string
	Transform() *Transform
	New() IEntity
}

type IClassname interface {
	Classname() string
}