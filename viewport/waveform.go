package viewport

import (
	_ "embed"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/glitter/shaders/waveform"
	"github.com/ignite-laboratories/host/graphics"
	"github.com/ignite-laboratories/host/graphics/math"
	"log"
	"time"
)

type Waveform[TValue core.Numeric] struct {
	*graphics.RenderableWindow
	*temporal.Dimension[TValue, any]

	// TimeScale represents the dimensional bounds to render the waveform within.
	TimeScale *std.TimeScale[TValue]

	// IsSigned indicates whether the type of this dimension is signed or not.
	IsSigned bool
}

func NewWaveform[TValue core.Numeric](windowSize std.XY[int], timeScale *std.TimeScale[TValue], isSigned bool, target *temporal.Dimension[TValue, any]) *Waveform[TValue] {
	v := &Waveform[TValue]{}
	v.TimeScale = timeScale
	v.RenderableWindow = graphics.SparkRenderableWindow(windowSize, v)
	v.Dimension = target
	v.IsSigned = isSigned
	return v
}

func (v *Waveform[TValue]) Initialize() {
	waveform.Init()
}

func (v *Waveform[TValue]) Render() {
	now := time.Now()
	oldest := now.Add(-v.TimeScale.Duration)
	v.Mutex.Lock()
	data := make([]std.Data[TValue], len(v.Timeline))
	copy(data, v.Timeline)
	v.Mutex.Unlock()

	// Clear the screen
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	locOfProjectionUniform := gl.GetUniformLocation(waveform.ProgramID, gl.Str("uProjectionMatrix\x00"))
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
	if v.IsSigned {
		projection = math.Ortho(0.0, float64(v.TimeScale.Duration), float64(-(v.TimeScale.Height / 2)), float64(v.TimeScale.Height/2), -1.0, 1.0) // Example
	} else {
		projection = math.Ortho(0.0, float64(v.TimeScale.Duration), 0, float64(v.TimeScale.Height), -1.0, 1.0) // Example
	}
	gl.UniformMatrix4fv(locOfProjectionUniform, 1, false, &projection[0])

	// Send them to the GPU using a VBO
	vbo := graphics.CreateVBO(vertices)

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
	gl.UseProgram(waveform.ProgramID)

	// Draw the line
	pointCount := len(vertices) / 2
	gl.DrawArrays(gl.LINE_STRIP, 0, int32(pointCount))

	// Cleanup
	gl.BindVertexArray(0)
	gl.DeleteBuffers(1, &vbo)
	gl.DeleteVertexArrays(1, &vao)
}
