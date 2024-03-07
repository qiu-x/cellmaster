package backends

import "cellmaster/gui/backends/gles"

type RenderingBackend interface {
	Init()
	RenderLoop()
	CleanUp()
}

func GetBestBackend() RenderingBackend {
	return &gles.GlesRenderer{}
}
