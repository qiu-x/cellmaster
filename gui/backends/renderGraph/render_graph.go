package renderGraph

import "cellmaster/gui"

type IRenderNode interface {
	Parent() *IRenderNode
	Children() *[]IRenderNode
	Load(gui.IContainer)
	Render()
}

type RenderNodeBase struct {
	parent   *IRenderNode
	children []IRenderNode
}

func (n *RenderNodeBase) Children() *[]IRenderNode {
	return &n.children
}

func (n *RenderNodeBase) Parent() *IRenderNode {
	return n.parent
}

type RenderNodeRoot struct {
	children []IRenderNode
}

func (n *RenderNodeRoot) Children() *[]IRenderNode {
	return &n.children
}

func (n *RenderNodeRoot) Load(gui.IContainer) {
	panic("unimplemented")
}

func (n *RenderNodeRoot) Parent() *IRenderNode {
	return nil
}

func (n *RenderNodeRoot) Render() {}

type RenderGraph struct {
	Root RenderNodeRoot
}

func NewRenderGraph() *RenderGraph {
	return &RenderGraph{
		Root: RenderNodeRoot{
			children: make([]IRenderNode, 0),
		},
	}
}
