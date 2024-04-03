package main

import (
	"cellmaster/gui"
	"cellmaster/gui/backends"
	_ "cellmaster/gui/elements"
	_ "cellmaster/gui/layouts"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	renderer := backends.GetBestBackend()
	defer renderer.CleanUp()

	// tilled := layouts.NewTiledLayout()
		// AddChild(&elements.Placeholder{}).
		// AddChild(&elements.Placeholder{})

	scene := gui.NewScene()
	// scene.Tree.AddChild(tilled)
	renderer.RenderLoop(scene)
}
