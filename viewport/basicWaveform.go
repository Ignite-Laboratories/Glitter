package viewport

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/glitter/assets"
	"github.com/ignite-laboratories/glitter/math"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type BasicWaveform[TValue core.Numeric] struct {
	*sdl2.Head

	// Dimension provides the temporal data that drives this waveform.
	Dimension *temporal.Dimension[TValue, any]

	// TimeScale represents the dimensional bounds to render the waveform within.
	TimeScale *std.TimeScale[TValue]

	// IsSigned indicates whether the type of this dimension is signed or not.
	IsSigned bool

	fragmentShader uint32
	vertexShader   uint32
	program        uint32
}

func NewBasicWaveform[TValue core.Numeric](fullscreen bool, framePotential core.Potential, title string, size *std.XY[int], pos *std.XY[int], timeScale *std.TimeScale[TValue], isSigned bool, target *temporal.Dimension[TValue, any]) *BasicWaveform[TValue] {
	view := &BasicWaveform[TValue]{}
	if fullscreen {
		view.Head = sdl2.CreateFullscreenWindow(core.Impulse, title, view, framePotential, false)
	} else {
		view.Head = sdl2.CreateWindow(core.Impulse, title, size, pos, view, framePotential, false)
	}
	view.TimeScale = timeScale
	view.Dimension = target
	view.IsSigned = isSigned

	view.EventHandler = view.TestInput

	return view
}

func (view *BasicWaveform[TValue]) TestInput(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		if e.Type == sdl.KEYDOWN {
			switch e.Keysym.Sym {
			case sdl.K_2:
				fmt.Println("here")
			}
		}
	}
}

func (view *BasicWaveform[TValue]) Impulse(ctx core.Context) {
	now := time.Now()
	oldest := now.Add(-view.TimeScale.Duration)
	view.Dimension.Mutex.Lock()
	data := make([]std.Data[TValue], len(view.Dimension.Timeline))
	copy(data, view.Dimension.Timeline)
	view.Dimension.Mutex.Unlock()

	// Clear the screen
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	locOfProjectionUniform := gl.GetUniformLocation(view.program, gl.Str("uProjectionMatrix\x00"))
	if locOfProjectionUniform == -1 {
		log.Fatalln("Failed to find uniform uProjectionMatrix")
	}

	// Prepare the vertices
	vertices := make([]float32, len(data)*2) // 2 floats per point (X, Y)
	var i int
	for _, d := range data {
		vertices[i] = float32(d.Moment.Sub(oldest))
		i++
		vertices[i] = float32(d.Point)
		i++
	}

	var projection []float32
	if view.IsSigned {
		projection = math.Ortho(0.0, float64(view.TimeScale.Duration), float64(-(view.TimeScale.Height / 2)), float64(view.TimeScale.Height/2), -1.0, 1.0) // Example
	} else {
		projection = math.Ortho(0.0, float64(view.TimeScale.Duration), 0, float64(view.TimeScale.Height), -1.0, 1.0) // Example
	}
	gl.UniformMatrix4fv(locOfProjectionUniform, 1, false, &projection[0])

	// Send them to the GPU using a VBO
	vbo := glitter.CreateVBO(vertices)

	// Set up the VAO
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	// Bind the VBO to the VAO
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)

	// Tell GL how to walk the vertex data (2 floats per point, 4 bytes per)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 2*4, nil)
	gl.EnableVertexAttribArray(0)

	// Draw the line
	gl.LineWidth(5.0)
	pointCount := len(vertices) / 2
	gl.DrawArrays(gl.LINE_STRIP, 0, int32(pointCount))

	// Cleanup
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (view *BasicWaveform[TValue]) Cleanup() {
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
}

func (view *BasicWaveform[TValue]) Initialize() {
	view.vertexShader = glitter.CompileShader(assets.Get.Shader("basicWaveform/basicWaveform.vert"), gl.VERTEX_SHADER)
	view.fragmentShader = glitter.CompileShader(assets.Get.Shader("basicWaveform/basicWaveform.frag"), gl.FRAGMENT_SHADER)
	view.program = glitter.LinkPrograms(view.vertexShader, view.fragmentShader)
}
