package viewport

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/glitter/assets"
	"github.com/ignite-laboratories/hydra"
	"log"
	"math"
	"sync"
)

type StackedByteWave[THeadDef any] struct {
	*hydra.Head[THeadDef]

	data []struct {
		Color std.RGBA[byte]
		Bytes []byte
	}
	bgColor std.RGBA[byte]

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

func NewStackedByteWave[THeadDef any](head *hydra.Head[THeadDef], bgColor std.RGBA[byte]) *StackedByteWave[THeadDef] {
	view := &StackedByteWave[THeadDef]{}
	view.bgColor = bgColor
	view.data = make([]struct {
		Color std.RGBA[byte]
		Bytes []byte
	}, 0)

	view.Head = head
	view.Head.SetImpulsable(view)

	return view
}

func (view *StackedByteWave[THeadDef]) SetBGColor(color std.RGBA[byte]) {
	view.Lock()
	view.bgColor = color
	view.Unlock()
}

func (view *StackedByteWave[THeadDef]) AddBytes(color std.RGBA[byte], bytes []byte) int {
	view.Lock()
	view.data = append(view.data, struct {
		Color std.RGBA[byte]
		Bytes []byte
	}{color, bytes})
	i := len(view.data) - 1
	view.Unlock()
	return i
}

func (view *StackedByteWave[THeadDef]) UpdateBytes(index int, color std.RGBA[byte], bytes []byte) {
	view.Lock()
	view.data[index] = struct {
		Color std.RGBA[byte]
		Bytes []byte
	}{color, bytes}
	view.Unlock()
}

func (view *StackedByteWave[THeadDef]) Lock() {
	view.mutex.Lock()
}

func (view *StackedByteWave[THeadDef]) Unlock() {
	view.mutex.Unlock()
}

func (view *StackedByteWave[THeadDef]) Initialize() {
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

func (view *StackedByteWave[THeadDef]) Impulse(ctx core.Context) {
	gl.ClearColor(view.bgColor.NormalizeToFloat32().RGBA())
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	for _, data := range view.data {
		view.drawWave(data.Color, data.Bytes)
	}
}

func (view *StackedByteWave[THeadDef]) drawWave(color std.RGBA[byte], data []byte) {
	if len(data) == 0 {
		return
	}

	r, g, b, a := color.NormalizeToFloat32().RGBA()
	gl.Uniform4f(view.colorLoc, r, g, b, a)

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

func (view *StackedByteWave[THeadDef]) Cleanup() {
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
	gl.DeleteVertexArrays(1, &view.vao)
	gl.DeleteBuffers(1, &view.vbo)
}
