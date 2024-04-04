package elements

import "cellmaster/gui"

// Placeholder element
// Fills the containing elements
type Placeholder struct {
	parent   *gui.IContainer
	children []gui.Slot
}

func (p *Placeholder) Slots() *[]gui.Slot {
	return &p.children
}

func (p *Placeholder) Parent() *gui.IContainer {
	return p.parent
}
