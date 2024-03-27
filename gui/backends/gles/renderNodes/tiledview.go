package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/gles/utils"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/layouts"

	gl "github.com/go-gl/gl/v3.1/gles2"
)

var VertexShader = `
#version 300 es
precision highp float;
in vec3 vert;

void main() {
    gl_Position = vec4(vert, 1);
}
` + "\x00"

var FragmentShader = `
#version 300 es
precision highp float;

uniform sampler2D tex;
in highp vec2 fragTexCoord;
out vec4 outputColor;

void main() {
    outputColor = vec4(0.1,0.1,0.1, 1);
}
` + "\x00"

type TiledView struct {
	renderGraph.RenderNodeBase
	reference *layouts.TiledView
	shader    uint32
	vao       uint32
}

var bgVerts = []float32{
	-1.0, 1.0, 0.1, 0.0, 0.0,
	-1.0, -1.0, 0.1, 0.0, 1.0,
	1.0, 1.0, 0.1, 1.0, 0.0,
	1.0, -1.0, 0.1, 1.0, 1.0,
}

func (t *TiledView) Load(tiledview gui.IContainer) {
	t.reference = tiledview.(*layouts.TiledView)

	var err error
	t.shader, err = utils.NewProgram(VertexShader, FragmentShader)
	if err != nil {
		panic(err)
	}

	gl.GenVertexArrays(1, &t.vao)
	gl.BindVertexArray(t.vao)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(bgVerts)*4, gl.Ptr(bgVerts), gl.STATIC_DRAW)

	vertAttrib := uint32(gl.GetAttribLocation(t.shader, gl.Str("vert\x00")))
	gl.EnableVertexAttribArray(vertAttrib)
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)
}

func (t *TiledView) Render() {
	gl.UseProgram(t.shader)
	gl.BindVertexArray(t.vao)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
