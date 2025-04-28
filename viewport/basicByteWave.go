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

type BasicByteWave struct {
	*sdl2.Head

	bytes   []byte
	mutex   sync.Mutex
	bgColor std.RGBA[float32]
	fgColor std.RGBA[float32]

	fragmentShader uint32
	vertexShader   uint32
	program        uint32
	vao            uint32
	vbo            uint32
	vertices       []float32
}

func NewBasicByteWave(engine *core.Engine, fullscreen bool, framePotential core.Potential, title string, size *std.XY[int], pos *std.XY[int], bgColor std.RGBA[byte], fgColor std.RGBA[byte], bytes []byte) *BasicByteWave {
	view := &BasicByteWave{}
	view.bytes = bytes
	view.bgColor = bgColor.NormalizeToFloat32()
	view.fgColor = fgColor.NormalizeToFloat32()

	if fullscreen {
		view.Head = sdl2.CreateFullscreenWindow(engine, title, view, framePotential, false)
	} else {
		view.Head = sdl2.CreateWindow(engine, title, size, pos, view, framePotential, false)
	}

	return view
}

func (view *BasicByteWave) SetBytes(bytes []byte) {
	view.mutex.Lock()
	view.bytes = bytes
	view.mutex.Unlock()
}

func (view *BasicByteWave) Lock() {
	view.mutex.Lock()
}

func (view *BasicByteWave) Unlock() {
	view.mutex.Unlock()
}

func (view *BasicByteWave) Initialize() {
	view.vertexShader = glitter.CompileShader(assets.Get.Shader("waveform/basicWaveform.vert"), gl.VERTEX_SHADER)
	view.fragmentShader = glitter.CompileShader(assets.Get.Shader("waveform/basicWaveformColor.frag"), gl.FRAGMENT_SHADER)
	view.program = glitter.LinkPrograms(view.vertexShader, view.fragmentShader)

	gl.UseProgram(view.program)

	gl.GenVertexArrays(1, &view.vao)
	gl.BindVertexArray(view.vao)

	gl.GenBuffers(1, &view.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, view.vbo)
}

func (view *BasicByteWave) Impulse(ctx core.Context) {
	data := make([]byte, len(view.bytes))
	copy(data, view.bytes)

	// Clear the screen
	gl.ClearColor(view.bgColor.RGBA())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	if len(data) == 0 {
		return
	}

	locOfProjectionUniform := gl.GetUniformLocation(view.program, gl.Str("uProjectionMatrix\x00"))
	if locOfProjectionUniform == -1 {
		log.Fatalln("Failed to find uniform 'uProjectionMatrix'")
	}

	colorLocation := gl.GetUniformLocation(view.program, gl.Str("fgColor\x00"))
	if colorLocation == -1 {
		log.Fatalln("Unable to find uniform location for 'fgColor'")
	}
	r, g, b, a := view.fgColor.RGBA()
	gl.Uniform4f(colorLocation, r, g, b, a)

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
	gl.UniformMatrix4fv(locOfProjectionUniform, 1, false, &projection[0])

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

func (view *BasicByteWave) Cleanup() {
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
	gl.DeleteVertexArrays(1, &view.vao)
	gl.DeleteBuffers(1, &view.vbo)
}
