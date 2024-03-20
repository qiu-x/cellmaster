package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/renderGraph"
)

type TiledView struct {
	renderGraph.RenderNodeBase
}

func (p *TiledView) From(gui.IContainer) *renderGraph.IRenderNode {
	panic("unimplemented")
}

func (p *TiledView) Render() {
	panic("unimplemented")
}
