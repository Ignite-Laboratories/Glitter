package viewport

import (
	_ "embed"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/host/sdl2"
	"time"
)

//go:embed shaders/tearing.frag
var fragmentShaderSource string

//go:embed shaders/tearing.vert
var vertexShaderSource string

type Tearing struct {
	*sdl2.Window

	fragmentShader uint32
	vertexShader   uint32
	program        uint32
	vao            uint32
	vbo            uint32
	vertices       []float32
}

func NewTearing(fullscreen bool, framePotential core.Potential, title string, size *std.XY[int], pos *std.XY[int]) *Tearing {
	view := &Tearing{}
	if fullscreen {
		view.Window = sdl2.CreateFullscreenWindow(core.Impulse, title, view, framePotential, false)
	} else {
		view.Window = sdl2.CreateWindow(core.Impulse, title, size, pos, view, framePotential, false)
	}

	return view
}

func (view *Tearing) Initialize() {
	view.vertexShader = glitter.CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	view.fragmentShader = glitter.CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	view.program = glitter.LinkPrograms(view.vertexShader, view.fragmentShader)

	gl.UseProgram(view.program)

	view.vertices = []float32{
		-1.0, -1.0, // Bottom-left corner
		1.0, -1.0, // Bottom-right corner
		-1.0, 1.0, // Top-left corner
		-1.0, 1.0, // Top-left corner
		1.0, -1.0, // Bottom-right corner
		1.0, 1.0, // Top-right corner
	}

	gl.GenVertexArrays(1, &view.vao)
	gl.GenBuffers(1, &view.vbo)

	gl.BindVertexArray(view.vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, view.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(view.vertices)*4, gl.Ptr(view.vertices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 0, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (view *Tearing) Impulse(ctx core.Context) {
	gl.ClearColor(0.25, 0.25, 0.25, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	screenWidth, screenHeight := view.Window.Handle.GetSize()
	resolutionUniform := gl.GetUniformLocation(view.program, gl.Str("resolution\x00"))
	gl.Uniform2f(resolutionUniform, float32(screenWidth), float32(screenHeight))

	elapsed := float32(time.Since(core.Inception).Seconds())
	timeUniform := gl.GetUniformLocation(view.program, gl.Str("time\x00"))
	gl.Uniform1f(timeUniform, elapsed)

	// Bind the VAO and draw the fullscreen quad
	gl.BindVertexArray(view.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.BindVertexArray(0)

	view.Window.Handle.GLSwap()
}

func (view *Tearing) Cleanup() {
	gl.DeleteVertexArrays(1, &view.vao)
	gl.DeleteBuffers(1, &view.vbo)
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
}
