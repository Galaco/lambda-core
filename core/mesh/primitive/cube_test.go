package primitive

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"reflect"
	"testing"
)

func TestNewCube(t *testing.T) {
	sut := NewCube()
	if reflect.TypeOf(sut) != reflect.TypeOf(&Cube{}) {
		t.Error("unexpected value returned when creating Cube")
	}
}

func TestCube_GetFaceMode(t *testing.T) {
	sut := NewCube()
	if sut.GetFaceMode() != gl.TRIANGLES {
		t.Errorf("unexpected face mode for Cube. Expected %d, but received: %d", gl.TRIANGLES, sut.GetFaceMode())
	}
}
