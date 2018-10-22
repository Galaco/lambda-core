package scene

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/entity"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/scene/loader"
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
}

func loadWorld(file *bsplib.Bsp) {
	currentScene.SetWorld(loader.LoadMap(file))
}

func loadEntities(entdata *lumps.EntData) {
	vmfEntityTree, err := loader.ParseEntities(entdata.GetData())
	if err != nil {
		debug.Fatal(err)
	}
	entityList := entitylib.FromVmfNodeTree(vmfEntityTree.Unclassified)
	debug.Logf("Found %d entities\n", entityList.Length())
	for i := 0; i < entityList.Length(); i++ {
		currentScene.AddEntity(loader.CreateEntity(entityList.Get(i)))
	}
}

func loadCamera() {
	currentScene.AddCamera(entity.NewCamera())
}
