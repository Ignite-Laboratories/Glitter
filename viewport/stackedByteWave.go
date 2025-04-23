package viewport

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/glitter/assets"
	"github.com/ignite-laboratories/hydra/sdl2"
	"log"
	"math"
	"sync"
)

type StackedByteWave struct {
	*sdl2.Head

	data []struct {
		Color std.RGBA
		Bytes []byte
	}
	bgColor std.RGBA

	mutex sync.Mutex

	fragmentShader uint32
	vertexShader   uint32
	program        uint32
	vao            uint32
	vbo            uint32
	projLoc        int32
	colorLoc       int32
	vertices       []float32
}

func NewStackedByteWave(engine *core.Engine, fullscreen bool, framePotential core.Potential, title string, size *std.XY[int], pos *std.XY[int], bgColor std.RGBA) *StackedByteWave {
	view := &StackedByteWave{}
	view.bgColor = bgColor
	view.data = make([]struct {
		Color std.RGBA
		Bytes []byte
	}, 0)

	if fullscreen {
		view.Head = sdl2.CreateFullscreenWindow(engine, title, view, framePotential, false)
	} else {
		view.Head = sdl2.CreateWindow(engine, title, size, pos, view, framePotential, false)
	}

	return view
}

func (view *StackedByteWave) SetBGColor(color std.RGBA) {
	view.Lock()
	view.bgColor = color
	view.Unlock()
}

func (view *StackedByteWave) AddBytes(color std.RGBA, bytes []byte) {
	view.Lock()
	view.data = append(view.data, struct {
		Color std.RGBA
		Bytes []byte
	}{color, bytes})
	view.Unlock()
}

func (view *StackedByteWave) Lock() {
	view.mutex.Lock()
}

func (view *StackedByteWave) Unlock() {
	view.mutex.Unlock()
}

func (view *StackedByteWave) Initialize() {
	view.vertexShader = glitter.CompileShader(assets.Get.Shader("waveform/basicWaveform.vert"), gl.VERTEX_SHADER)
	view.fragmentShader = glitter.CompileShader(assets.Get.Shader("waveform/basicWaveformColor.frag"), gl.FRAGMENT_SHADER)
	view.program = glitter.LinkPrograms(view.vertexShader, view.fragmentShader)

	gl.UseProgram(view.program)

	gl.GenVertexArrays(1, &view.vao)
	gl.BindVertexArray(view.vao)

	gl.GenBuffers(1, &view.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, view.vbo)

	view.projLoc = gl.GetUniformLocation(view.program, gl.Str("uProjectionMatrix\x00"))
	if view.projLoc == -1 {
		log.Fatalln("Failed to find uniform 'uProjectionMatrix'")
	}

	view.colorLoc = gl.GetUniformLocation(view.program, gl.Str("fgColor\x00"))
	if view.colorLoc == -1 {
		log.Fatalln("Unable to find uniform location for 'fgColor'")
	}
}

func (view *StackedByteWave) Impulse(ctx core.Context) {
	gl.ClearColor(view.bgColor.RGBA())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, data := range view.data {
		view.drawWave(data.Color, data.Bytes)
	}
}

func (view *StackedByteWave) drawWave(color std.RGBA, data []byte) {
	if len(data) == 0 {
		return
	}

	gl.Uniform4f(color.SplitRGBAWithLocation(view.colorLoc))

	// Prepare the vertices
	vertices := make([]float32, len(data)*2) // 2 floats per point (X, Y)
	var i int
	for x, d := range data {
		vertices[i] = float32(x)
		i++
		vertices[i] = float32(d)
		i++
	}

	var projection = glitter.Ortho(0.0, float64(len(data)-1), 0, float64(math.MaxUint8), -1.0, 1.0)
	gl.UniformMatrix4fv(view.projLoc, 1, false, &projection[0])

	// Send the vertices to the GPU using a VBO
	if len(vertices) > 0 {
		gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
	}

	// Tell GL how to walk the vertex data (2 floats per point, 4 bytes per)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 2*4, nil)
	gl.EnableVertexAttribArray(0)

	// Draw the line
	gl.LineWidth(2.5)
	pointCount := len(vertices) / 2
	gl.DrawArrays(gl.LINE_STRIP, 0, int32(pointCount))
}

func (view *StackedByteWave) Cleanup() {
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
	gl.DeleteVertexArrays(1, &view.vao)
	gl.DeleteBuffers(1, &view.vbo)
}
