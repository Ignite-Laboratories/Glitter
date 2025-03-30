package glitter

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
)

type Viewport struct {
	*temporal.Dimension[core.Runtime, chan core.Context]
	Title   *string
	Size    *std.XY[int]
	Window  *glfw.Window
	impulse chan core.Context
}

func NewViewport(engine *core.Engine, potential core.Potential, title *string, windowSize *std.XY[int]) *Viewport {
	// Spark off a new open GL context on an impulse thread
	v := &Viewport{}
	v.Title = title
	v.Size = windowSize
	v.impulse = make(chan core.Context)
	v.Dimension = temporal.DedicatedLoop(engine, potential, false, v.tick, v.cleanup)
	return v
}

func (v *Viewport) cleanup() {
	glfw.Terminate()
}

func (v *Viewport) createWindow() {
	InitializeGLFW()
	var err error

	v.Window, err = glfw.CreateWindow(v.Size.X, v.Size.Y, *v.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	v.Window.MakeContextCurrent()

	InitializeGL()
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func (v *Viewport) tick(ctx core.Context) {
	if v.Window == nil {
		v.createWindow()
	}

	if v.Window.ShouldClose() {
		v.Dimension.Destroy()
		return
	}
	ctx.Beat = 0

	glfw.PollEvents()
	v.Window.SwapBuffers()
	fmt.Println("tick")
}
