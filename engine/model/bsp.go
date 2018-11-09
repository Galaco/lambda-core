package model

import "github.com/galaco/Gource-Engine/engine/mesh"

type Bsp struct {
	internalMesh mesh.IMesh
	faces []mesh.Face
	visibleFaces []*mesh.Face
	numFaces int32
}

func (bsp *Bsp) Mesh() mesh.IMesh {
	return bsp.internalMesh
}

func (bsp *Bsp) Faces() []mesh.Face {
	return bsp.faces[:bsp.numFaces]
}

func (bsp *Bsp) VisibleFaces() []*mesh.Face {
	return bsp.visibleFaces
}

func (bsp *Bsp) AddFace(face *mesh.Face) {
	bsp.faces[bsp.numFaces] = *face
	bsp.numFaces++
}

func (bsp *Bsp) SetFaces(faces []mesh.Face) {
	bsp.faces = faces
	bsp.numFaces = int32(len(faces))
}

func (bsp *Bsp) SetVisibleFaces(faces []*mesh.Face) {
	bsp.visibleFaces = faces
}

func NewBsp(refMesh *mesh.Mesh) *Bsp {
	return &Bsp{
		internalMesh: refMesh,
		faces: make([]mesh.Face, 65536),
		numFaces: 0,
	}
}
