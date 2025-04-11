package viewport

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/glitter/math"
	"github.com/ignite-laboratories/glitter/shaders/waveform"
	"github.com/ignite-laboratories/host/hydra"
	"log"
	"time"
)

type Waveform[TValue core.Numeric] struct {
	*hydra.Head
	*temporal.Dimension[TValue, any]

	// TimeScale represents the dimensional bounds to render the waveform within.
	TimeScale *std.TimeScale[TValue]

	// IsSigned indicates whether the type of this dimension is signed or not.
	IsSigned bool
}

func NewWaveform[TValue core.Numeric](fullscreen bool, framePotential core.Potential, title string, size *std.XY[int], pos *std.XY[int], timeScale *std.TimeScale[TValue], isSigned bool, target *temporal.Dimension[TValue, any]) *Waveform[TValue] {
	wave := &Waveform[TValue]{}
	if fullscreen {
		wave.Head = hydra.CreateFullscreenWindow(core.Impulse, title, wave.Initialize, wave.Render, framePotential, false)
	} else {
		wave.Head = hydra.CreateWindow(core.Impulse, title, size, pos, wave.Initialize, wave.Render, framePotential, false)
	}
	wave.TimeScale = timeScale
	wave.Dimension = target
	wave.IsSigned = isSigned

	return wave
}

func (w *Waveform[TValue]) Render(ctx core.Context) {
	now := time.Now()
	oldest := now.Add(-w.TimeScale.Duration)
	w.Mutex.Lock()
	data := make([]std.Data[TValue], len(w.Timeline))
	copy(data, w.Timeline)
	w.Mutex.Unlock()

	// Clear the screen
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	locOfProjectionUniform := gl.GetUniformLocation(waveform.SimpleProgram, gl.Str("uProjectionMatrix\x00"))
	if locOfProjectionUniform == -1 {
		log.Fatalln("Failed to find uniform uProjectionMatrix")
	}

	// Prepare the vertices
	vertices := make([]float32, len(data)*2) // 2 floats per point (X, Y)
	var i int
	for _, d := range data {
		vertices[i] = float32(d.Moment.Sub(oldest).Seconds()) // Convert to seconds
		i++
		vertices[i] = float32(d.Point)
		i++
	}

	var projection []float32
	if w.IsSigned {
		projection = math.Ortho(0.0, w.TimeScale.Duration.Seconds(), float64(-(w.TimeScale.Height / 2)), float64(w.TimeScale.Height/2), -1.0, 1.0) // Example
	} else {
		projection = math.Ortho(0.0, w.TimeScale.Duration.Seconds(), 0, float64(w.TimeScale.Height), -1.0, 1.0) // Example
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

	// Tell GL which shader program to use
	gl.UseProgram(waveform.SimpleProgram)

	// Draw the line
	//gl.LineWidth(5.0)
	pointCount := len(vertices) / 2
	gl.DrawArrays(gl.LINE_STRIP, 0, int32(pointCount))

	// Cleanup
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}

func (w *Waveform[TValue]) Initialize() {
	waveform.Init()
}
