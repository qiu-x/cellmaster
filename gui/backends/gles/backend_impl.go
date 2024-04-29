package gles

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles/renderNodes"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/elements"
	"cellmaster/gui/layouts"
	"cellmaster/gui/scenegraph"
	"fmt"

	gl "github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GlesRenderer struct {
	glesWindow *glfw.Window
	window gui.Window
}

// GetWindow implements backends.RenderingBackend.
func (r *GlesRenderer) GetWindow() *gui.Window {
	return &r.window
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
	glfw.WindowHint(glfw.Resizable, glfw.True)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	r.glesWindow, err = glfw.CreateWindow(1280, 720, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}

	r.glesWindow.MakeContextCurrent()

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

func (r *GlesRenderer) RenderLoop(scene *scenegraph.Scene) {
	for !r.glesWindow.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update scene size
		// TODO: Add if resized check
		width, height := r.glesWindow.GetFramebufferSize()
		r.window.Size.Height = width
		r.window.Size.Height = height

		renderGraph := ParseScene(&scene.Root)
		RenderTree(&renderGraph.Root)
		r.glesWindow.SwapBuffers()
		glfw.PollEvents()
		gl.Viewport(0, 0, int32(width), int32(height))
	}
}

func (r *GlesRenderer) CleanUp() {
	glfw.Terminate()
}

func getRenderNode(v gui.IElement) renderGraph.IRenderNode {
	switch v.(type) {
	case gui.IContainer:
		return getContainerRenderer(v)
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

func getContainerRenderer(e gui.IElement) renderGraph.IRenderNode {
	v := e.(gui.IContainer)
	switch v.(type) {
	case *layouts.TiledView:
		c := &renderNodes.TiledView{}
		c.Load(v)
		return c
	default:
		return &renderNodes.Placeholder{}
	}
}

func ParseScene(scene *scenegraph.SceneRoot) *renderGraph.RenderGraph {
	rg := renderGraph.NewRenderGraph()
	var copyTree func(gui.IContainer, renderGraph.IRenderNode)
	copyTree = func(scn gui.IContainer, prev renderGraph.IRenderNode) {
		if scn == nil {
			return
		}
		for _, v := range *scn.Slots() {
			elem := v.Element
			current := getRenderNode(elem)
			*prev.Children() = append(*prev.Children(), current)
			if container, ok := (elem).(gui.IContainer); ok {
				copyTree(container, current)
			}
		}
	}

	// Set root RenderNode
	mainContainer, ok := scene.MainView.AsContainer()
	if !ok {
		return rg
	}

	rootRenderNode := getRenderNode(mainContainer)
	*rg.Root.Children() = append(*rg.Root.Children(), rootRenderNode)
	copyTree(mainContainer, rootRenderNode)

	return rg
}
