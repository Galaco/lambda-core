package scene

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/loader"
	entity2 "github.com/galaco/Gource-Engine/engine/loader/entity"
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	entitylib "github.com/galaco/source-tools-common/entity"
)

func LoadFromFile(fileName string) {
	bspData, err := bsplib.ReadFromFile(fileName)
	if err != nil {
		debug.Fatal(err)
	}
	if bspData.GetHeader().Version < 20 {
		debug.Fatal("Unsupported BSP Version. Exiting...")
	}

	//Set pakfile for filesystem
	filesystem.SetPakfile(bspData.GetLump(bsplib.LUMP_PAKFILE).(*lumps.Pakfile))

	loadWorld(bspData)

	loadEntities(bspData.GetLump(bsplib.LUMP_ENTITIES).(*lumps.EntData))

	loadCamera()

	currentScene.isLoaded = true
}

func loadWorld(file *bsplib.Bsp) {
	currentScene.SetWorld(loader.LoadMap(file))
}

func loadEntities(entdata *lumps.EntData) {
	vmfEntityTree, err := entity2.ParseEntities(entdata.GetData())
	if err != nil {
		debug.Fatal(err)
	}
	entityList := entitylib.FromVmfNodeTree(vmfEntityTree.Unclassified)
	debug.Logf("Found %d entities\n", entityList.Length())
	for i := 0; i < entityList.Length(); i++ {
		currentScene.AddEntity(entity2.CreateEntity(entityList.Get(i)))
	}

	skyCamera := entityList.FindByKeyValue("classname", "sky_camera")
	if skyCamera == nil {
		return
	}

	worldSpawn := entityList.FindByKeyValue("classname", "worldspawn")
	if worldSpawn == nil {
		return
	}

	currentScene.world.BuildSkybox(loader.LoadSky(worldSpawn.ValueForKey("skyname")), skyCamera.VectorForKey("origin"), float32(skyCamera.IntForKey("scale")))
}

func loadCamera() {
	currentScene.AddCamera(entity.NewCamera())
}
