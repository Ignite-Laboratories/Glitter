package viewport

import (
	_ "embed"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/glitter/assets"
	"github.com/ignite-laboratories/hydra"
	"github.com/ignite-laboratories/hydra/glfw"
	"github.com/ignite-laboratories/hydra/sdl2"
	"sync"
	"time"
)

type ScreenTearTester[THeadDef any] struct {
	*hydra.Head[THeadDef]

	fragmentShader uint32
	vertexShader   uint32
	program        uint32
	vao            uint32
	vbo            uint32
	vertices       []float32
	mutex          sync.Mutex

	getSize func() (int, int)
}

func NewSDL2ScreenTearTester(head *hydra.Head[sdl2.SDLDefinition]) *ScreenTearTester[sdl2.SDLDefinition] {
	view := &ScreenTearTester[sdl2.SDLDefinition]{}
	view.Head = head
	view.Head.SetImpulsable(view)
	view.getSize = func() (int, int) {
		x, y := view.Head.Definition.Handle.GetSize()
		return int(x), int(y)
	}

	return view
}

func NewGLFWScreenTearTester(head *hydra.Head[glfw.GLFWDefinition]) *ScreenTearTester[glfw.GLFWDefinition] {
	view := &ScreenTearTester[glfw.GLFWDefinition]{}
	view.Head = head
	view.Head.SetImpulsable(view)
	view.getSize = view.Head.Definition.Handle.GetSize

	return view
}

func (view *ScreenTearTester[THeadDef]) Lock() {
	view.mutex.Lock()
}

func (view *ScreenTearTester[THeadDef]) Unlock() {
	view.mutex.Unlock()
}

func (view *ScreenTearTester[THeadDef]) Initialize() {
	view.vertexShader = glitter.CompileShader(assets.Get.Shader("screenTearTester/screenTearTester.vert"), gl.VERTEX_SHADER)
	view.fragmentShader = glitter.CompileShader(assets.Get.Shader("screenTearTester/screenTearTester.frag"), gl.FRAGMENT_SHADER)
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

func (view *ScreenTearTester[THeadDef]) Impulse(ctx core.Context) {
	gl.ClearColor(0.25, 0.25, 0.25, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	screenWidth, screenHeight := view.getSize()
	resolutionUniform := gl.GetUniformLocation(view.program, gl.Str("resolution\x00"))
	gl.Uniform2f(resolutionUniform, float32(screenWidth), float32(screenHeight))

	elapsed := float32(time.Since(core.Inception).Seconds())
	timeUniform := gl.GetUniformLocation(view.program, gl.Str("time\x00"))
	gl.Uniform1f(timeUniform, elapsed)

	// Bind the VAO and draw the fullscreen quad
	gl.BindVertexArray(view.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.BindVertexArray(0)
}

func (view *ScreenTearTester[THeadDef]) Cleanup() {
	gl.DeleteVertexArrays(1, &view.vao)
	gl.DeleteBuffers(1, &view.vbo)
	gl.DeleteShader(view.vertexShader)
	gl.DeleteShader(view.fragmentShader)
	gl.DeleteProgram(view.program)
}
