package gui

type IContainer interface {
	Parent() *IContainer
	Slots() *[]ISlot
}

type IElement interface {
	Parent() *IContainer
}

type ISlot interface {
	Dimetions() Rect
	Element() IElement
}

type Rect struct {
	X, Y, Width, Height int
}

type SceneGraph struct {
	parent   *IContainer
	children []ISlot
}

func (sg *SceneGraph) Slots() *[]ISlot {
	return &sg.children
}

func (sg *SceneGraph) Parent() *IContainer {
	return sg.parent
}

func (sg *SceneGraph) AddChild(node ISlot) {
	sg.children = append(sg.children, node)
}

type Scene struct {
	Tree SceneGraph
}

func NewScene() *Scene {
	tree := SceneGraph{
		parent:   nil,
		children: make([]ISlot, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
