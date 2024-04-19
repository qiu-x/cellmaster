package scenegraph

import "cellmaster/gui"

type SceneRoot struct {
	MainView gui.Slot
}

type Scene struct {
	Root SceneRoot
	window gui.Window
}

func NewScene(elem gui.IElement, win gui.Window) *Scene {
	root := SceneRoot{
		MainView: gui.Slot{
			Dimetions: gui.Rect{}, // value not used
			Element: elem,
		},
	}
	scene := &Scene{
		Root: root,
		window: win,
	}
	return scene
}
