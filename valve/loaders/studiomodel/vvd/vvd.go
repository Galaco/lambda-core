package vvd

import "github.com/go-gl/mathgl/mgl32"

const MAX_NUM_LODS = 8
const MAX_NUM_BONES_PER_VERT = 3

type Vvd struct {
	Header   header
	Fixups   []fixup
	Vertices []vertex
	Tangents []mgl32.Vec4
}

type header struct {
	Id               int32
	Version          int32
	Checksum         int32
	NumLODs          int32
	NumLODVertexes   [MAX_NUM_LODS]int32
	NumFixups        int32
	FixupTableStart  int32
	VertexDataStart  int32
	TangentDataStart int32
}

type fixup struct {
	Lod            int32
	SourceVertexID int32
	NumVertexes    int32
}

type vertex struct {
	BoneWeight boneWeight
	Position   mgl32.Vec3
	Normal     mgl32.Vec3
	UVs        mgl32.Vec2
}

type boneWeight struct {
	Weight   [MAX_NUM_BONES_PER_VERT]float32
	Bone     [MAX_NUM_BONES_PER_VERT]int8
	NumBones int8
}
