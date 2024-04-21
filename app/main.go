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
	renderer := backends.GetBestBackend()
	window := renderer.GetWindow()
	defer renderer.CleanUp()

	tilled := layouts.NewTiledLayout().
		AddChild(&elements.Placeholder{}).
		AddChild(&elements.Placeholder{})

	scene := scenegraph.NewScene(tilled, window)
	renderer.RenderLoop(scene)
}
