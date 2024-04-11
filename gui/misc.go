package gui

type Rect struct {
	X, Y, Width, Height int
}

type IContainer interface {
	Parent() *Slot
	Slots() *[]Slot
}

type IElement interface {
	Parent() *Slot
}

type Slot struct {
	Dimetions Rect
	Element   IElement
}

func (s* Slot) AsContainer() (cnt IContainer, ok bool) {
	cnt, ok = s.Element.(IContainer)
	return
}

