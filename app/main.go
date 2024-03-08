package main

import (
	"cellmaster/gui"
	"cellmaster/gui/backends"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	renderer := backends.GetBestBackend()
	defer renderer.CleanUp()

	scene := gui.NewScene()
	renderer.RenderLoop(scene)
}
