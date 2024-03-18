package layouts

import (
	"cellmaster/gui"
)

type TiledView struct {
	parent   *gui.IContainer
	children []*gui.IElement
	sizes    []float64
	offsets  []float64
}

func (t TiledView) Children() *[]*gui.IElement {
	return &t.children
}

func (t TiledView) Parent() *gui.IContainer {
	return t.parent
}

func (t *TiledView) AddChild(child gui.IElement) {
	t.children = append(t.children, &child)
}

func NewTiledLayout(count int) *TiledView {
	tiledView := &TiledView{
		children: make([]*gui.IElement, 0),
		sizes:    make([]float64, count),
		offsets:  make([]float64, count),
	}
	tiledView.Recompute()
	return tiledView
}

func (gl *TiledView) Recompute() {
	for i := range gl.sizes {
		gl.sizes[i] = 1/float64(len(gl.children)) + gl.offsets[i]
	}
}

func (gl *TiledView) Resize(row int, newSize float64) {
	if row > len(gl.children)-2 {
		panic("Resized row does not exist or is not resizable (ie. last row)")
	}
	if newSize > 1 {
		newSize = 1
	}
	for i := range gl.sizes {
		gl.offsets[i] -= newSize / float64(len(gl.children)-1)
	}
	gl.offsets[row] = newSize
	gl.Recompute()
}
