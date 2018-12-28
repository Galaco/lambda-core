package primitive

import (
	"github.com/galaco/Gource-Engine/gosigl"
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

func TestCube_GetFaceMode(t *testing.T) {
	t.Skip("cannot instantiate without providing an opengl context first")
	sut := NewCube()
	if sut.GetFaceMode() != gosigl.Triangles {
		t.Errorf("unexpected face mode for Cube. Expected %d, but received: %d", gosigl.Triangles, sut.GetFaceMode())
	}
}
