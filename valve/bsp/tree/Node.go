package tree

type INode interface {
	IsLeaf() bool
}

type Node struct {
	Children [2]INode
}

func (node *Node) IsLeaf() bool {
	return false
}
func (node *Node) AddChild(index int, child INode) {
	node.Children[index] = child
}


type Leaf struct {
	FaceIndexList []uint16
}

func (leaf *Leaf) IsLeaf() bool {
	return true
}
