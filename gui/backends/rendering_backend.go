package backends

import (
	"cellmaster/gui"
	"cellmaster/gui/elements"
	"cellmaster/gui/backends/gles"
)

type RenderingImplementation interface {
	GetPlaceholderRenderer(elements.Placeholder)
}

type RenderingBackend interface {
	Init()
	RenderLoop(*gui.Scene)
	CleanUp()
}

func GetBestBackend() RenderingBackend {
	renderer := &gles.GlesRenderer{}
	renderer.Init()
	return renderer
}
