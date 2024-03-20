package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/layouts"
)

type TiledView struct {
	renderGraph.RenderNodeBase
	reference layouts.TiledView
}

func (t *TiledView) Load(tiledview gui.IContainer) {
	t.reference = tiledview.(layouts.TiledView)
}

func (t *TiledView) Render() {
	panic("unimplemented")
}
