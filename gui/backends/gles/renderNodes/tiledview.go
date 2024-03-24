package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/layouts"

	"github.com/go-gl/gl/v2.1/gl"
)

var VertexShader = `
#version 330
uniform vec2 window_size;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;

void main() {
	mat4 window_scale = mat4 (
		vec4(window_size.x/window_size.y, 0.0, 0.0, 0.0),
		vec4(0.0, 1.0, 0.0, 0.0),
		vec4(0.0, 0.0, 1.0, 0.0),
		vec4(0.0, 0.0, 0.0, 1.0)
	);
	mat4 transform = mat4 (
		vec4(1.0, 0.0, 0.0, 0.0),
		vec4(0.0, 1.0, 0.0,-0.95),
		vec4(0.0, 0.0, 1.0, 0.0),
		vec4(0.0, 0.0, 0.0, 1.0)
	);
    fragTexCoord = vertTexCoord;
    gl_Position = vec4(vert, 1) * transform * window_scale;
}
` + "\x00"

var FragmentShader = `
#version 330
uniform sampler2D tex;
in vec2 fragTexCoord;
out vec4 outputColor;
void main() {
	outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

type TiledView struct {
	renderGraph.RenderNodeBase
	reference layouts.TiledView
}

var bgVerts = []float32{
	-1.0, 1.0, 0.1, 0.0, 0.0,
	-1.0, -1.0, 0.1, 0.0, 1.0,
	1.0, 1.0, 0.1, 1.0, 0.0,
	1.0, -1.0, 0.1, 1.0, 1.0,
}

func (t *TiledView) Load(tiledview gui.IContainer) {
	t.reference = tiledview.(layouts.TiledView)

	program, err := newProgram(VertexShader, FragmentShader)
	if err != nil {
		panic(err)
	}

	var vao uint32
	gl.GenVertexArrays(1, &text_vao)
	gl.BindVertexArray(text_vao)

	var vbo uint32
	gl.GenBuffers(1, &text_vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, text_vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(bgVerts)*4, gl.Ptr(bgVerts), gl.STATIC_DRAW)
}

func (t *TiledView) Render() {
	panic("unimplemented")
}
