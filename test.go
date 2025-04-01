package glitter

import (
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
)

type Red struct {
	*Viewport
}

func CreateRed(title string, size std.XY[int]) *Red {
	t := &Red{}
	t.Viewport = CreateViewport(title, size, t.Render)
	return t
}

func (v *Red) Render(ctx core.Context) {
	// Set the clear color (red in this case)
	gl.ClearColor(1.0, 0.0, 0.0, 1.0) // RGBA: Red
	gl.Clear(gl.COLOR_BUFFER_BIT)     // Clear the screen with the set color

	sdl.GL_SwapWindow(v.window)
}

type Blue struct {
	*Viewport
}

func CreateBlue(title string, size std.XY[int]) *Blue {
	t := &Blue{}
	t.Viewport = CreateViewport(title, size, t.Render)
	return t
}

func (v *Blue) Render(ctx core.Context) {
	// Set the clear color (red in this case)
	gl.ClearColor(0.0, 0.0, 1.0, 1.0) // RGBA: Red
	gl.Clear(gl.COLOR_BUFFER_BIT)     // Clear the screen with the set color

	sdl.GL_SwapWindow(v.window)
}
