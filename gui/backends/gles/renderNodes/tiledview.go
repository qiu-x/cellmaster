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
uniform vec4 borderColor;

void main() {
    gl_Position = vec4(vert, 1);
}
` + "\x00"

var FragmentShader = `
#version 300 es
precision highp float;

uniform vec4 Color;
out vec4 outputColor;

void main() {
    outputColor = Color;
}
` + "\x00"

type TiledView struct {
	renderGraph.RenderNodeBase
	reference  *layouts.TiledView
	shader     uint32
	vao        uint32
	borderColor [4]float32 // RGBA color for the border
	innerColor [4]float32 // RGBA color for the rects
}

func (t *TiledView) Load(tiledview gui.IContainer) {
	t.reference = tiledview.(*layouts.TiledView)
	t.borderColor = [4]float32{0.7, 0.7, 0.7, 0}
	t.innerColor = [4]float32{0.1, 0.1, 0.1, 0}

	var err error
	t.shader, err = utils.NewProgram(VertexShader, FragmentShader)
	if err != nil {
		panic(err)
	}

	gl.GenVertexArrays(1, &t.vao)
	gl.BindVertexArray(t.vao)
}

func (t *TiledView) Render() {
	gl.UseProgram(t.shader)
	gl.BindVertexArray(t.vao)

	// Draw 4 rectangles next to each other
	numRectangles := 4
	rectWidth := 1.0 / float32(numRectangles)


	for i := 0; i < numRectangles; i++ {
		// // Set the color for the border
		gl.Uniform4fv(gl.GetUniformLocation(t.shader, gl.Str("Color\x00")), 1, &t.innerColor[0])

		offset := float32(i) * rectWidth
		// Define the vertices for the current rectangle
		verts := []float32{
			-1.0 + offset, 1.0, 0.1, 0.0, 0.0,
			-1.0 + offset, -1.0, 0.1, 0.0, 1.0,
			-1.0 + offset + rectWidth, 1.0, 0.1, 1.0, 0.0,
			-1.0 + offset + rectWidth, -1.0, 0.1, 1.0, 1.0,
		}

		var vbo uint32
		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(verts)*4, gl.Ptr(verts), gl.STATIC_DRAW)

		vertAttrib := uint32(gl.GetAttribLocation(t.shader, gl.Str("vert\x00")))
		gl.EnableVertexAttribArray(vertAttrib)
		gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)

		// Draw the current rectangle
		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)

		// Set the color for the border
		gl.Uniform4fv(gl.GetUniformLocation(t.shader, gl.Str("Color\x00")), 1, &t.borderColor[0])

		// Draw the borders between rectangles
		if i < numRectangles-1 {
			borderVerts := []float32{
				-1.0 + offset + rectWidth, 1.0, 0.1, 0.0, 0.0,
				-1.0 + offset + rectWidth, -1.0, 0.1, 0.0, 1.0,
				-1.0 + offset + rectWidth + 0.02, 1.0, 0.1, 1.0, 0.0, // Adjust border width as needed
				-1.0 + offset + rectWidth + 0.02, -1.0, 0.1, 1.0, 1.0, // Adjust border width as needed
			}

			var borderVBO uint32
			gl.GenBuffers(1, &borderVBO)
			gl.BindBuffer(gl.ARRAY_BUFFER, borderVBO)
			gl.BufferData(gl.ARRAY_BUFFER, len(borderVerts)*4, gl.Ptr(borderVerts), gl.STATIC_DRAW)

			gl.VertexAttribPointerWithOffset(vertAttrib, 3, gl.FLOAT, false, 5*4, 0)

			gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
		}
	}
}
