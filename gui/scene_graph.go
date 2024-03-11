package gui

type INode interface {
	Parent() INode
	Children() []INode
}

type Node struct {
	parent   INode
	children []INode
}

func (n *Node) Children() []INode {
	return n.children
}

func (n *Node) Parent() INode {
	return n.parent
}

type SceneGraph struct {
	Node
	parent   INode
	children []INode
}

func (st *SceneGraph) AddChild(node INode) {
	st.children = append(st.children, node)
}

type Scene struct {
	Tree SceneGraph
}

func NewScene() *Scene {
	tree := SceneGraph{
		parent:   nil,
		children: make([]INode, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
