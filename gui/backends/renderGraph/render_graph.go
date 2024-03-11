package renderGraph

import "cellmaster/gui"

type IRenderNode interface {
	Parent() *IRenderNode
	Children() []*IRenderNode
	Render()
}

type RenderNodeBase struct {
	parent   *IRenderNode
	children []*IRenderNode
}

func (n *RenderNodeBase) Children() []*IRenderNode {
	return n.children
}

func (n *RenderNodeBase) Parent() *IRenderNode {
	return n.parent
}

type RenderGraph struct {
	RenderNodeBase
	parent   *IRenderNode
	children []*IRenderNode
}

func NewRenderGraph() *RenderGraph {
	return &RenderGraph{
		parent:   nil,
		children: make([]*IRenderNode, 0),
	}
}

func (n *RenderGraph) FromSceneGraph(sceneGraph *gui.SceneGraph) *RenderGraph {
	var copyTree func(node gui.INode, tree *RenderGraph)
	copyTree = func(node gui.INode, tree *RenderGraph) {
		if node == nil {
			return
	    }
		for _, v := range node.Children() {
			copyTree(v, tree)
		}
	}
    copyTree(sceneGraph, n)
    return n
}

