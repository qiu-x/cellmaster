package backends

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
