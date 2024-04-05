package elements

import "cellmaster/gui"

// Placeholder element
// Fills the containing elements
type Placeholder struct {
	parent   *gui.Slot
}

func (p *Placeholder) Parent() *gui.Slot {
	return p.parent
}
