package core

// ETypes exist as a means of "reflection" without employing it properly, as its a bit slow.
// Etypes are link to types
type EType string

func (t *EType) String() string {
	return string(*t)
}
