package gles

import (
	"cellmaster/gui"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GlesRenderer struct {
	window *glfw.Window
}

func (r *GlesRenderer) Init() {
	var err error
	err = glfw.Init()
	if err != nil {
		panic(err)
	}

	r.window, err = glfw.CreateWindow(640, 480, "Testing", nil, nil)
	if err != nil {
		panic(err)
	}
	r.window.MakeContextCurrent()
}

func (r *GlesRenderer) RenderLoop(scene *gui.Scene) {
	for !r.window.ShouldClose() {
		r.window.SwapBuffers()
		glfw.PollEvents()
	}
}

func (r *GlesRenderer) CleanUp() {
	glfw.Terminate()
}
