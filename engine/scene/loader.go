package scene

import (
	"github.com/galaco/Gource-Engine/engine/core/debug"
	"github.com/galaco/Gource-Engine/engine/factory"
	"github.com/galaco/Gource-Engine/engine/filesystem"
	"github.com/galaco/Gource-Engine/engine/scene/loader"
	entity2 "github.com/galaco/Gource-Engine/entity"
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	"github.com/galaco/source-tools-common/entity"
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

	loadSky(currentScene.GetWorld().KeyValues().ValueForKey("skyname"))
}

func loadWorld(file *bsplib.Bsp) {
	worldSpawn := factory.NewEntity(loader.LoadMap(file)).(*entity2.WorldSpawn)
	currentScene.SetWorld(worldSpawn)
}

func loadEntities(entdata *lumps.EntData) {
	vmfEntityTree, err := loader.ParseEntities(entdata.GetData())
	if err != nil {
		debug.Fatal(err)
	}
	entityList := entity.FromVmfNodeTree(vmfEntityTree.Unclassified)
	debug.Logf("Found %d entities\n", entityList.Length())
	for i := 0; i < entityList.Length(); i++ {
		//	currentScene.AddEntity(bsp.CreateEntity(entityList.Get(i)))
	}

	currentScene.GetWorld().SetKeyValues(entityList.FindByKeyValue("classname", "worldspawn"))
}

func loadSky(skyname string) {
	currentScene.GetWorld()

	sky, err := loader.LoadSky(skyname)
	if err == nil {
		factory.NewComponent(sky, currentScene.GetWorld())
	}

}
