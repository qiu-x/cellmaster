package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/renderGraph"
)

type Placeholder struct {
	renderGraph.RenderNodeBase
}

func (p *Placeholder) Load(gui.IContainer) {}

func (p *Placeholder) Render() {}
