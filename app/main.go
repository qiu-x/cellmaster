package main

import (
	"cellmaster/gui/backends"
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

	scene := scenegraph.NewScene(nil, window)
	renderer.RenderLoop(scene)
}
