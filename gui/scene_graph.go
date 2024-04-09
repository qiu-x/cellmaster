package gui

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

type Rect struct {
	X, Y, Width, Height int
}

type SceneRoot struct {
	Dimetions *Rect
	MainView   IContainer
}

type Scene struct {
	Root SceneRoot
	WindowRect Rect
}

func NewScene(elem IContainer, rect *Rect) *Scene {
	root := SceneRoot{
		Dimetions: rect,
		MainView:  elem,
	}
	scene := &Scene{
		Root: root,
	}
	return scene
}
