package renderNodes

import (
	"cellmaster/gui"
	"cellmaster/gui/backends/renderGraph"
	"cellmaster/gui/layouts"
	"fmt"
	"strings"

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
    outputColor = vec4(0,0,1, 1);
}
` + "\x00"

type TiledView struct {
	renderGraph.RenderNodeBase
	reference *layouts.TiledView
	shader uint32
	vao uint32
}

var bgVerts = []float32{
	-1.0, 1.0, 0.1, 0.0, 0.0,
	-1.0, -1.0, 0.1, 0.0, 1.0,
	1.0, 1.0, 0.1, 1.0, 0.0,
	1.0, -1.0, 0.1, 1.0, 1.0,
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func newProgram(vertexShaderSource, fragmentShaderSource string) (uint32, error) {
	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return program, nil
}

func (t *TiledView) Load(tiledview gui.IContainer) {
	t.reference = tiledview.(*layouts.TiledView)

	var err error
	t.shader, err = newProgram(VertexShader, FragmentShader)
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
	gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)}

func (t *TiledView) Render() {
	gl.UseProgram(t.shader)
	gl.BindVertexArray(t.vao)
	gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
}
