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
	tree := SceneRoot{
		parent: nil,
		MainView: Slot{
			Dimetions: Rect{
				X:      0,
				Y:      0,
				Width:  600,
				Height: 400,
			},
			Element: elem,
		},
	}
	scene := &Scene{
		Root: tree,
	}
	return scene
}
