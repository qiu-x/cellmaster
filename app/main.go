package main

import (
	"cellmaster/gui"
	"cellmaster/gui/backends"
	"cellmaster/gui/layouts"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	renderer := backends.GetBestBackend()
	defer renderer.CleanUp()

	scene := gui.NewScene()
	scene.Tree.AddChild(layouts.NewGridLayout(2, 2))
	renderer.RenderLoop(scene)
}
