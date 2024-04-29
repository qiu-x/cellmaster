package gio

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles/renderNodes"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/elements"
	"cellmaster/gui/layouts"
	"cellmaster/gui/scenegraph"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/event"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type GioRenderer struct {
	appWindow *app.Window
	window gui.Window
}

// GetWindow implements backends.RenderingBackend.
func (r *GioRenderer) GetWindow() *gui.Window {
	return &r.window
}

func (r *GioRenderer) Init() {
	r.appWindow.Option(app.Size(unit.Dp(800), unit.Dp(700)))
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

func (r *GioRenderer) RenderLoop(scene *scenegraph.Scene) {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	events := make(chan event.Event)
	acks := make(chan struct{})

	go func() {
		for {
			ev := r.appWindow.Event()
			events <- ev
			<-acks
			if _, ok := ev.(app.DestroyEvent); ok {
				return
			}
		}
	}()

	var ops op.Ops
	for {
		select {
		case e := <-events:
			switch e := e.(type) {
			case app.DestroyEvent:
				acks <- struct{}{}
				return
			case app.FrameEvent:
				gtx := app.NewContext(&ops, e)
				e.Frame(gtx.Ops)
			}
			acks <- struct{}{}
		}
	}
}

func (r *GioRenderer) CleanUp() {
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
