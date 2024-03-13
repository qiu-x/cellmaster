package renderGraph

type IRenderNode interface {
	Parent() *IRenderNode
	Children() *[]*IRenderNode
	Render()
}

type RenderNodeRoot struct {
	children []*IRenderNode
}

func (n *RenderNodeRoot) Children() *[]*IRenderNode {
	return &n.children
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
			children: make([]*IRenderNode, 0),
		},
	}
}
