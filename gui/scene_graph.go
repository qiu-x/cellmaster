package gui

type IContainer interface {
	Parent() *IContainer
	Slots() *[]Slot
}

type IElement interface {
	Parent() *IContainer
}

type Slot struct {
	Dimetions Rect
	Element IElement
}

type Rect struct {
	X, Y, Width, Height int
}

type SceneGraph struct {
	parent   *IContainer
	children []Slot
}

func (sg *SceneGraph) Slots() *[]Slot {
	return &sg.children
}

func (sg *SceneGraph) Parent() *IContainer {
	return sg.parent
}

func (sg *SceneGraph) AddChild(node Slot) {
	sg.children = append(sg.children, node)
}

type Scene struct {
	Tree SceneGraph
}

func NewScene() *Scene {
	tree := SceneGraph{
		parent:   nil,
		children: make([]Slot, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
