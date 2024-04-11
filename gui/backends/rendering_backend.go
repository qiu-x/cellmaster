package backends

import (
	"cellmaster/gui/backends/gles"
	"cellmaster/gui/scenegraph"
)

type RenderingBackend interface {
	Init()
	RenderLoop(*scenegraph.Scene)
	CleanUp()
}

func GetBestBackend() RenderingBackend {
	renderer := &gles.GlesRenderer{}
	renderer.Init()
	return renderer
}
