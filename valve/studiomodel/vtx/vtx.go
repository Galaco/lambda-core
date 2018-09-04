package vtx

type Vtx struct {
	Header header
	BodyParts []bodyPartHeader
	Models []modelHeader
	ModelLODs []modelLODHeader
	Meshes []meshHeader
	StripGroups []stripGroupHeader
	Indices []uint16
	Vertices []float32
	Strips []stripHeader
	VertexList []vertex
}


type header struct {
	Version int32
	VertCacheSize int32

	MaxBonesperStrip uint16
	MaxBonesPerTriangle uint16
	MaxBonesPerVert int32

	CheckSum int32
	NumLODs int32

	MaterialReplacementListOffset int32

	NumBodyParts int32
	BodyPartOffset int32
}


type bodyPartHeader struct {
	NumModels int32
	ModelOffset int32
}

type modelHeader struct {
	NumLODs int32
	LODOffset int32
}

type modelLODHeader struct {
	NumMeshes int32
	MeshOffset int32
	SwitchPoint float32
}


type meshHeader struct {
	NumStripGroups int32
	StripGroupHeaderOffset int32

	Flags uint8
}

const StripGroupIsFlexed = 0x01
const StripGroupIsHWSkinned = 0x02
const StripGroupIsDeltaFlexed = 0x04
const StripGroupSuppressHWMorph = 0x08

type stripGroupHeader struct {
	NumVerts int32
	VertOffset int32

	NumIndices int32
	IndexOffset int32

	NumStrips int32
	StripOffset int32

	Flags uint8
}


type stripHeader struct {
	NumIndices int32
	IndexOffset int32

	NumVerts int32
	VertOffset int32

	NumBones int16

	Flags uint8

	NumBoneStateChanges int32
	BoneStateChangeOffset int32
}


type vertex struct {
	BoneWeightIndex [3]uint8
	NumBones uint8

	OriginalMeshVertexID uint16

	BoneID [3]int8
}