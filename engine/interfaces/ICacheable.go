package interfaces

import "github.com/galaco/go-me-engine/engine/core"

type ICacheable interface {
	GetHandle() core.Handle
	Update(float64)
}
