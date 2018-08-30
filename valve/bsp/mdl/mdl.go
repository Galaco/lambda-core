package mdl

import "github.com/go-gl/mathgl/mgl32"

type studiohdr struct {
	id         int32
	version    int32
	name       [64]byte // 64 char exactly, null byte padded
	dataLength int32

	eyeposition   mgl32.Vec3
	illumposition mgl32.Vec3
	hullMin       mgl32.Vec3
	hullMax       mgl32.Vec3
	viewBBMin     mgl32.Vec3
	viewBBMax     mgl32.Vec3

	flags int32

	//studio bone
	boneCount  int32
	boneOffset int32
	//studiobonecontroller
	boneControllerCount  int32
	boneControllerOffset int32
	//mstudiohitboxset
	hitboxCount  int32
	hitboxOffset int32
	//mstudioanimdesc
	localAnimationCount  int32
	localAnimationOffset int32
	//mstudioseqdesc
	localSequenceCount  int32
	localSequenceOffset int32

	activityListVersion int32
	eventsIndexed       int32

	//vmt filenames - mstudiotexture
	textureCount  int32
	textureOffset int32

	textureDirCount  int32
	textureDirOffset int32

	skinReferenceCount       int32
	skinReferenceFamilyCount int32
	skinReferenceIndex       int32

	// mstudiobodyparts
	bodyPartCount  int32
	bodypartOffset int32

	// mstudioattachment
	attachmentCount  int32
	attachmentOffset int32

	localNodeCount     int32
	localNodeIndex     int32
	localNodeNameIndex int32

	// mstudioflexdesc
	flexDescCount int32
	flexDescIndex int32

	// mstudioflexcontroller
	flexControllerCount int32
	flexControllerIndex int32

	//mstudioflexrule
	flexRulesCount int32
	flexRulesIndex int32

	//mstudioikchain
	ikChainCount int32
	ikChainIndex int32

	//mstudiomouth
	mouthsCount int32
	mouthsIndex int32

	//mstudioposeparamdesc
	localPoseParamCount int32
	localPoseParamIndex int32

	surfacePropertyIndex int32

	keyValueIndex int32
	keyValueCount int32

	// mstudioiklock
	ikLockCount int32
	ikLockIndex int32

	mass     float32
	contents int32

	// mstudiomodelgroup
	includeModelCount int32
	includeModelIndex int32

	virtualModel int32

	// mstudianimblock
	animblocksNameIndex int32
	animblocksCount     int32
	animblocksIndex     int32

	animblockModel int32

	boneTableNameIndex int32

	vertexBase int32
	offsetBase int32

	directionalDotProduct byte
	rootLOD               uint8
	numAllowedRootLods    uint8

	_ byte
	_ int32

	flexControllerUICount int32
	flexControllerUIIndex int32

	// otional studiohdr2 offset
	studioHDR2Index int32

	_ int32
}

type Mdl struct {
	header studiohdr
}
