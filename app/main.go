package main

import (
	"cellmaster/gui/backends"
	"cellmaster/gui/elements"
	"cellmaster/gui/layouts"
	"cellmaster/gui/scenegraph"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	// TODO: Separate window out of renderers
	renderer := backends.GetBestBackend()
	defer renderer.CleanUp()

	tilled := layouts.NewTiledLayout().
		AddChild(&elements.Placeholder{}).
		AddChild(&elements.Placeholder{})

	scene := scenegraph.NewScene(tilled)
	renderer.RenderLoop(scene)
}
