package interfaces

// Generic game manager.
// Different systems should implement these methods
type IManager interface {
	Register()
	Update(float64)
	Unregister()
}
