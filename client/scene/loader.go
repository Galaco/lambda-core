package scene

import (
	"github.com/galaco/Gource-Engine/client/config"
	"github.com/galaco/Gource-Engine/client/scene/visibility"
	"github.com/galaco/Gource-Engine/client/scene/world"
	"github.com/galaco/Gource-Engine/core/entity"
	"github.com/galaco/Gource-Engine/core/filesystem"
	"github.com/galaco/Gource-Engine/core/loader"
	entity2 "github.com/galaco/Gource-Engine/core/loader/entity"
	"github.com/galaco/Gource-Engine/core/logger"
	"github.com/galaco/Gource-Engine/core/model"
	bsplib "github.com/galaco/bsp"
	"github.com/galaco/bsp/lumps"
	entitylib "github.com/galaco/source-tools-common/entity"
	"github.com/go-gl/mathgl/mgl32"
)

func LoadFromFile(fileName string) {
	bspData, err := bsplib.ReadFromFile(fileName)
	if err != nil {
		logger.Fatal(err)
	}
	if bspData.GetHeader().Version < 20 {
		logger.Fatal("Unsupported BSP Version. Exiting...")
	}

	//Set pakfile for filesystem
	filesystem.RegisterPakfile(bspData.GetLump(bsplib.LUMP_PAKFILE).(*lumps.Pakfile))

	loadWorld(bspData)

	loadEntities(bspData.GetLump(bsplib.LUMP_ENTITIES).(*lumps.EntData))

	loadCamera()

	currentScene.isLoaded = true
}

func loadWorld(file *bsplib.Bsp) {
	baseWorld := loader.LoadMap(file)

	baseWorldBsp := baseWorld.Bsp()
	baseWorldBspFaces := baseWorldBsp.ClusterLeafs()[0].Faces
	baseWorldStaticProps := baseWorld.StaticProps()

	visData := visibility.NewVisFromBSP(file)
	visLump := file.GetLump(bsplib.LUMP_VISIBILITY).(*lumps.Visibility).GetData()
	bspClusters := make([]model.ClusterLeaf, visLump.NumClusters)
	defaultCluster := model.ClusterLeaf{
		Id: 32767,
	}
	for _, bspLeaf := range visData.Leafs {
		for _, leafFace := range visData.LeafFaces[bspLeaf.FirstLeafFace : bspLeaf.FirstLeafFace+bspLeaf.NumLeafFaces] {
			if bspLeaf.Cluster == -1 {
				//defaultCluster.Faces = append(defaultCluster.Faces, bspFaces[leafFace])
				continue
			}
			bspClusters[bspLeaf.Cluster].Id = bspLeaf.Cluster
			bspClusters[bspLeaf.Cluster].Faces = append(bspClusters[bspLeaf.Cluster].Faces, baseWorldBspFaces[leafFace])
		}
	}

	// Assign staticprops to clusters
	for idx, prop := range baseWorld.StaticProps() {
		for _, leafId := range prop.LeafList() {
			clusterId := visData.Leafs[leafId].Cluster
			if clusterId == -1 {
				//defaultCluster.StaticProps = append(defaultCluster.StaticProps, &staticProps[idx])
				continue
			}
			bspClusters[clusterId].StaticProps = append(bspClusters[clusterId].StaticProps, &baseWorldStaticProps[idx])
		}
	}

	for _, idx := range baseWorldBsp.ClusterLeafs()[0].DispFaces {
		defaultCluster.Faces = append(defaultCluster.Faces, baseWorldBspFaces[idx])
	}

	baseWorldBsp.SetClusterLeafs(bspClusters)
	baseWorldBsp.SetDefaultCluster(defaultCluster)

	currentScene.SetWorld(world.NewWorld(*baseWorld.Bsp(), baseWorld.StaticProps(), visData))
}

func loadEntities(entdata *lumps.EntData) {
	vmfEntityTree, err := entity2.ParseEntities(entdata.GetData())
	if err != nil {
		logger.Fatal(err)
	}
	entityList := entitylib.FromVmfNodeTree(vmfEntityTree.Unclassified)
	logger.Notice("Found %d entities\n", entityList.Length())
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

	currentScene.world.BuildSkybox(
		loader.LoadSky(worldSpawn.ValueForKey("skyname")),
		skyCamera.VectorForKey("origin"),
		float32(skyCamera.IntForKey("scale")))
}

func loadCamera() {
	currentScene.AddCamera(entity.NewCamera(mgl32.DegToRad(70), float32(config.Get().Video.Width)/float32(config.Get().Video.Height)))
}
