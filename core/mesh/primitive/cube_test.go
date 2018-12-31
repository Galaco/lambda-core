package primitive

import (
	"reflect"
	"testing"
)

func TestNewCube(t *testing.T) {
	t.Skip("cannot instantiate without providing an opengl context first")
	sut := NewCube()
	if reflect.TypeOf(sut) != reflect.TypeOf(&Cube{}) {
		t.Error("unexpected value returned when creating Cube")
	}
}
