package renderNodes

import "cellmaster/gui/backends/renderGraph"

type Placeholder struct {
	renderGraph.RenderNodeBase
}

func (p *Placeholder) Render() {
	panic("unimplemented")
}
