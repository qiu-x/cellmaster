package gui

type IContainer interface {
	Parent() *IContainer
	Children() *[]*IElement
}

type IElement interface {
	Parent() *IContainer
}

type SceneGraph struct {
	parent   *IContainer
	children []*IElement
}

func (sg *SceneGraph) Children() *[]*IElement {
	return &sg.children
}

func (sg *SceneGraph) Parent() *IContainer {
	return sg.parent
}

func (sg *SceneGraph) AddChild(node IElement) {
	sg.children = append(sg.children, &node)
}

type Scene struct {
	Tree SceneGraph
}

func NewScene() *Scene {
	tree := SceneGraph{
		parent:   nil,
		children: make([]*IElement, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
