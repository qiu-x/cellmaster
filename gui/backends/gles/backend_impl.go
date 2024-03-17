package gles

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles/renderNodes"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/elements"
	"cellmaster/gui/layouts"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type GlesRenderer struct {
	window *glfw.Window
}

func (r *GlesRenderer) Init() {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	r.window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	r.window.MakeContextCurrent()
}

func (r *GlesRenderer) RenderLoop(scene *gui.Scene) {
	for !r.window.ShouldClose() {
		r.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (r *GlesRenderer) CleanUp() {
	glfw.Terminate()
}

func getRenderNode(v gui.IElement) renderGraph.IRenderNode {
	switch v.(type) {
	case gui.IElement:
		return getElementRenderer(v)
	case gui.IContainer:
		return getContainerRenderer(v.(gui.IContainer))
	default:
		return &renderNodes.Placeholder{}
	}
}

func getElementRenderer(v gui.IElement) renderGraph.IRenderNode {
	switch v.(type) {
	case *layouts.TiledView:
		return &renderNodes.Placeholder{}
	case *elements.Placeholder:
		return &renderNodes.Placeholder{}
	default:
		return &renderNodes.Placeholder{}
	}
}

func getContainerRenderer(v gui.IContainer) renderGraph.IRenderNode {
	switch v.(type) {
	case *layouts.TiledView:
		return &renderNodes.Placeholder{}
	default:
		return &renderNodes.Placeholder{}
	}
}

func ParseSceneGraph(sceneGraph *gui.SceneGraph) *renderGraph.RenderGraph {
	rg := renderGraph.NewRenderGraph()
	var copyTree func(node gui.IContainer, prev renderGraph.IRenderNode)
	copyTree = func(node gui.IContainer, prev renderGraph.IRenderNode) {
		if node == nil {
			return
		}
		for _, v := range *node.Children() {
			current := getRenderNode(*v)
			*current.Parent() = prev
			*prev.Children() = append(*current.Children(), &current)
			copyTree(*v, current)
		}
	}
	copyTree(sceneGraph, &rg.Root)
	return rg
}
