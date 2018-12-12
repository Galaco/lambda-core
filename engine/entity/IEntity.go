package entity

import (
	"github.com/galaco/source-tools-common/entity"
)

// IEntity Base interface
// All game entities need to implement this
type IEntity interface {
	KeyValues() *entity.Entity
	SetKeyValues(entity2 *entity.Entity)
	Classname() string
	Transform() *Transform
	New() IEntity
}

// IClassname all valid game entities should have a classname,
// but there may be temporary non-game entities that have classnames
type IClassname interface {
	Classname() string
}
