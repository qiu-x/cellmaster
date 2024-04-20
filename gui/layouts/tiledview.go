package layouts

import (
	"cellmaster/gui"
)

type TiledView struct {
	parent   *gui.Slot
	children []gui.Slot
	sizes    []float64
	offsets  []float64
}

func (t *TiledView) Slots() *[]gui.Slot {
	return &t.children
}

func (t *TiledView) Parent() *gui.Slot {
	return t.parent
}

func (t *TiledView) AddChild(child gui.IElement) *TiledView {
	t.children = append(t.children, gui.Slot{
		Dimetions: gui.Rect{},
		Element:   child,
	})
	t.sizes = append(t.sizes, 0)
	t.offsets = append(t.offsets, 0)
	t.ComputeLayout()
	return t
}

func NewTiledLayout() *TiledView {
	tiledView := &TiledView{
		children: make([]gui.Slot, 0),
		sizes:    make([]float64, 0),
		offsets:  make([]float64, 0),
	}
	tiledView.ComputeLayout()
	return tiledView
}

func (t *TiledView) ComputeLayout() {
	for i := range t.sizes {
		t.sizes[i] = 1/float64(len(t.children)) + t.offsets[i]
	}
}

func (t *TiledView) Resize(elm int, newSize float64) {
	if elm > len(t.children)-2 {
		panic("Resized row does not exist or is not resizable (ie. last row)")
	}
	if newSize > 1 {
		newSize = 1
	}
	for i := range t.sizes {
		t.offsets[i] -= newSize / float64(len(t.children)-1)
	}
	t.offsets[elm] = newSize
	t.ComputeLayout()
}
