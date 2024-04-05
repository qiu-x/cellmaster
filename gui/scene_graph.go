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
	children []Slot
}

func (sg *SceneRoot) Slots() *[]Slot {
	return &sg.children
}

func (sg *SceneRoot) Parent() *Slot {
	return nil
}

func (sg *SceneRoot) AddChild(node IElement) {
	sg.children = append(sg.children, Slot{
		Dimetions: Rect{},
		Element:   node,
	})
}

type Scene struct {
	Tree SceneRoot
}

func NewScene() *Scene {
	tree := SceneRoot{
		parent:   nil,
		children: make([]Slot, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
