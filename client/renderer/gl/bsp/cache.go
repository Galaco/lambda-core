package bsp

import (
	"github.com/galaco/Gource-Engine/core/event"
	"github.com/galaco/Gource-Engine/core/resource/message"
	"github.com/galaco/gosigl"
)

var MapGPUResource *gosigl.VertexObject

func SyncMapToGpu(dispatched event.IMessage) {
	msg := dispatched.(*message.MapLoaded)
	mesh := msg.Resource.Mesh()
	MapGPUResource = gosigl.NewMesh(mesh.Vertices())
	gosigl.CreateVertexAttribute(MapGPUResource, mesh.TextureCoordinates(), 2)
	gosigl.CreateVertexAttribute(MapGPUResource, mesh.Normals(), 3)

	// @TODO Find a better solution
	if len(mesh.LightmapCoordinates()) < 2 {
		gosigl.CreateVertexAttribute(MapGPUResource, []float32{0, 1}, 2)
	} else {
		gosigl.CreateVertexAttribute(MapGPUResource, mesh.LightmapCoordinates(), 2)
	}
	gosigl.FinishMesh()
}
