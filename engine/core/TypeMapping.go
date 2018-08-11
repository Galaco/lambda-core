package core

type EType string

func (t *EType) String() string {
	return string(*t)
}
