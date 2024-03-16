package gui

type IContainer interface {
	Parent() *IContainer
	Children() *[]*IContainer
}

type IElement interface {
	Parent() *IContainer
}

type SceneGraph struct {
	parent   *IContainer
	children []*IContainer
}

func (sg *SceneGraph) Children() *[]*IContainer {
	return &sg.children
}

func (sg *SceneGraph) Parent() *IContainer {
	return sg.parent
}

func (sg *SceneGraph) AddChild(node IContainer) {
	sg.children = append(sg.children, &node)
}

type Scene struct {
	Tree SceneGraph
}

func NewScene() *Scene {
	tree := SceneGraph{
		parent:   nil,
		children: make([]*IContainer, 0),
	}
	scene := &Scene{
		Tree: tree,
	}
	return scene
}
