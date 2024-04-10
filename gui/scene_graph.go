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
	parent   *IContainer
	MainView Slot
}

type Scene struct {
	Root SceneRoot
}

func NewScene(elem IContainer) *Scene {
	root := SceneRoot{
		parent: nil,
		MainView: Slot{
			Dimetions: Rect{}, // value not used
			Element: elem,
		},
	}
	scene := &Scene{
		Root: root,
	}
	return scene
}
