package prop

import (
	"github.com/galaco/Gource-Engine/core/event"
	"github.com/galaco/Gource-Engine/core/resource/message"
	"github.com/galaco/gosigl"
)

var ModelIdMap map[string][]*gosigl.VertexObject


func SyncPropToGpu(dispatched event.IMessage) {
	msg := dispatched.(*message.PropLoaded)
	if ModelIdMap[msg.Resource.GetFilePath()] != nil {
		return
	}
	vals := make([]*gosigl.VertexObject, len(msg.Resource.GetMeshes()))

	for idx,mesh := range msg.Resource.GetMeshes() {
		gpuObject := gosigl.NewMesh(mesh.Vertices())
		gosigl.CreateVertexAttribute(gpuObject, mesh.TextureCoordinates(), 2)
		gosigl.CreateVertexAttribute(gpuObject, mesh.Normals(), 3)

		// @TODO Find a better solution
		if len(mesh.LightmapCoordinates()) < 2 {
			lightmapCoordinates := []float32{0, 1}
			gosigl.CreateVertexAttribute(gpuObject, lightmapCoordinates, 2)
		} else {
			gosigl.CreateVertexAttribute(gpuObject, mesh.LightmapCoordinates(), 2)
		}
		gosigl.FinishMesh()
		vals[idx] = gpuObject
	}
	ModelIdMap[msg.Resource.GetFilePath()] = vals
}

func DestroyPropOnGPU(dispatched event.IMessage) {
	msg := dispatched.(*message.PropUnloaded)
	for _,i := range ModelIdMap[msg.Resource.GetFilePath()] {
		gosigl.DeleteMesh(i)
	}
}