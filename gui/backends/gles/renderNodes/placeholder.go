package renderNodes

import "cellmaster/gui/backends/renderGraph"

// Placeholder renderNode
type Placeholder struct {
}

func (p *Placeholder) Children() *[]*renderGraph.IRenderNode {
	panic("unimplemented")
}

func (p *Placeholder) Parent() *renderGraph.IRenderNode {
	panic("unimplemented")
}

func (p *Placeholder) Render() {
	panic("unimplemented")
}
