package model

import "github.com/galaco/Gource-Engine/engine/mesh"

type Bsp struct {
	internalMesh mesh.IMesh

	clusterLeafs []ClusterLeaf
	visibleClusterLeafs []*ClusterLeaf
}

func (bsp *Bsp) Mesh() mesh.IMesh {
	return bsp.internalMesh
}

func (bsp *Bsp) ClusterLeafs() []ClusterLeaf {
	return bsp.clusterLeafs
}

func (bsp *Bsp) VisibleClusterLeafs() []*ClusterLeaf {
	return bsp.visibleClusterLeafs
}

func (bsp *Bsp) SetClusterLeafs(clusterLeafs []ClusterLeaf) {
	bsp.clusterLeafs = clusterLeafs
}

func (bsp *Bsp) SetVisibleClusters(clusterLeafs []*ClusterLeaf) {
	bsp.visibleClusterLeafs = clusterLeafs
}

func NewBsp(refMesh *mesh.Mesh) *Bsp {
	return &Bsp{
		internalMesh: refMesh,
	}
}
