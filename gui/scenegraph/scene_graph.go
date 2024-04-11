package scenegraph

import "cellmaster/gui"

type SceneRoot struct {
	MainView gui.Slot
}

type Scene struct {
	Root SceneRoot
}

func NewScene(elem gui.IElement) *Scene {
	root := SceneRoot{
		MainView: gui.Slot{
			Dimetions: gui.Rect{}, // value not used
			Element: elem,
		},
	}
	scene := &Scene{
		Root: root,
	}
	return scene
}
