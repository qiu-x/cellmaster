package gles

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles/renderNodes"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/elements"
	"cellmaster/gui/layouts"
	"fmt"

	gl "github.com/go-gl/gl/v3.1/gles2"
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

	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLAnyProfile)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	r.window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	r.window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)
}

func RenderTree(node renderGraph.IRenderNode) {
	node.Render()
	if node.Children() == nil {
		return
	}
	for _, v := range *(node.Children()) {
		RenderTree(v)
	}
}

func (r *GlesRenderer) RenderLoop(scene *gui.Scene) {
	for !r.window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		renderGraph := ParseSceneGraph(&scene.Tree)
		// renderGraph.Root.Render()
		RenderTree(&renderGraph.Root)
		r.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (r *GlesRenderer) CleanUp() {
	glfw.Terminate()
}

func getRenderNode(v gui.IElement) renderGraph.IRenderNode {
	switch v.(type) {
	case gui.IContainer:
		return getContainerRenderer(v.(gui.IContainer))
	case gui.IElement:
		return getElementRenderer(v)
	default:
		return &renderNodes.Placeholder{}
	}
}

func getElementRenderer(v gui.IElement) renderGraph.IRenderNode {
	switch v.(type) {
	case *elements.Placeholder:
		return &renderNodes.Placeholder{}
	default:
		return &renderNodes.Placeholder{}
	}
}

func getContainerRenderer(v gui.IContainer) renderGraph.IRenderNode {
	switch v.(type) {
	case *layouts.TiledView:
		c := &renderNodes.TiledView{}
		c.Load(v)
		return c
	default:
		return &renderNodes.Placeholder{}
	}
}

func ParseSceneGraph(sceneGraph *gui.SceneGraph) *renderGraph.RenderGraph {
	rg := renderGraph.NewRenderGraph()
	var copyTree func(gui.IContainer, renderGraph.IRenderNode)
	copyTree = func(scn gui.IContainer, prev renderGraph.IRenderNode) {
		if scn == nil {
			return
		}
		for _, v := range *scn.Slots() {
			elem := v.Element
			current := getRenderNode(elem)
			if prev.Parent() != nil {
				*current.Parent() = prev
			}
			*prev.Children() = append(*prev.Children(), current)
			if container, ok := (elem).(gui.IContainer); ok {
				copyTree(container, current)
			}
		}
	}
	copyTree(sceneGraph, &rg.Root)
	return rg
}
