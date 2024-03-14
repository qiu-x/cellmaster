package renderNodes

import "cellmaster/gui/backends/renderGraph"

// Placeholder renderNode
type Placeholder struct {
	renderGraph.RenderNodeBase
}

func (p *Placeholder) Render() {
	panic("unimplemented")
}
