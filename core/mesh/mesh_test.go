package mesh

import (
	"github.com/galaco/Gource-Engine/core/material"
	"github.com/galaco/Gource-Engine/core/texture"
	"reflect"
	"testing"
)

func TestNewMesh(t *testing.T) {
	if reflect.TypeOf(NewMesh()) != reflect.TypeOf(&Mesh{}) {
		t.Errorf("unexpected type returned for NewMesh. Expected: %s, but received: %s", reflect.TypeOf(&Mesh{}), reflect.TypeOf(NewMesh()))
	}
}

func TestMesh_AddLightmapCoordinate(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddLightmapCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.LightmapCoordinates()[i] != expected[i] {
			t.Error("unexpected lightmap coordinate")
		}
	}
}

func TestMesh_AddNormal(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddNormal(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Normals()[i] != expected[i] {
			t.Error("unexpected normal")
		}
	}
}

func TestMesh_AddTextureCoordinate(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddTextureCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.TextureCoordinates()[i] != expected[i] {
			t.Error("unexpected texture coordinate")
		}
	}
}

func TestMesh_AddVertex(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddVertex(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Vertices()[i] != expected[i] {
			t.Error("unexpected vertex")
		}
	}
}

func TestMesh_Bind(t *testing.T) {
	t.Skip()
}

func TestMesh_Destroy(t *testing.T) {
	t.Skip()
}

func TestMesh_Finish(t *testing.T) {
	t.Skip()
}

func TestMesh_GetLightmap(t *testing.T) {
	sut := Mesh{}
	expected := &texture.Lightmap{}
	sut.SetLightmap(expected)

	if expected != sut.GetLightmap() {
		t.Error("unexpected lightmap applied to mesh")
	}
}

func TestMesh_GetMaterial(t *testing.T) {
	sut := Mesh{}
	expected := &material.Material{
		FilePath: "foo.vmt",
	}
	sut.SetMaterial(expected)

	if expected != sut.GetMaterial() {
		t.Error("unexpected material applied to mesh")
	}
}

func TestMesh_LightmapCoordinates(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddLightmapCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.LightmapCoordinates()[i] != expected[i] {
			t.Error("unexpected lightmap coordinate")
		}
	}
}

func TestMesh_Normals(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddNormal(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Normals()[i] != expected[i] {
			t.Error("unexpected normal")
		}
	}
}

func TestMesh_SetLightmap(t *testing.T) {
	sut := Mesh{}
	expected := &texture.Lightmap{}
	sut.SetLightmap(expected)

	if expected != sut.GetLightmap() {
		t.Error("unexpected lightmap applied to mesh")
	}
}

func TestMesh_SetMaterial(t *testing.T) {
	sut := Mesh{}
	expected := &material.Material{
		FilePath: "foo.vmt",
	}
	sut.SetMaterial(expected)

	if expected != sut.GetMaterial() {
		t.Error("unexpected material applied to mesh")
	}
}

func TestMesh_TextureCoordinates(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddTextureCoordinate(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.TextureCoordinates()[i] != expected[i] {
			t.Error("unexpected texture coordinate")
		}
	}
}

func TestMesh_Vertices(t *testing.T) {
	sut := Mesh{}
	expected := []float32{
		1, 2, 3, 4,
	}
	sut.AddVertex(expected...)

	for i := 0; i < len(expected); i++ {
		if sut.Vertices()[i] != expected[i] {
			t.Error("unexpected vertex")
		}
	}
}
