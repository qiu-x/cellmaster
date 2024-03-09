package gui

import "cellmaster/gui"

type INode interface {
	Parent() *INode
	Children() []*INode
}

type Node struct {
	parent   *INode
	children []*INode
}

func (n *Node) Children() []*INode {
	return n.children
}

func (n *Node) Parent() *INode {
	return n.parent
}

type SceneTree struct {
	Node
	parent   *INode
	children []*INode
}

func (st *SceneTree) AddChild(node gui.INode) {
	st.children = append(st.children, node)
}

type Scene struct {
	Tree SceneTree
}

func NewScene() *Scene {
	tree := SceneTree{
		parent:   nil,
		children: make([]*INode, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
