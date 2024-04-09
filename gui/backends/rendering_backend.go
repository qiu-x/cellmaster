package backends

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles"
)

type RenderingBackend interface {
	Init()
	RenderLoop(*gui.Scene)
	CleanUp()
	WindowRectRef() *gui.Rect
}

func GetBestBackend() RenderingBackend {
	renderer := &gles.GlesRenderer{}
	renderer.Init()
	return renderer
}
