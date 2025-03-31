package glitter

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"log"
)

// factory is a means of creating neural activity from a single host thread for graphical purposes.
type factory struct {
	*temporal.Dimension[core.Runtime, chan core.Context]
	viewports   map[uint64]*Viewport
	initialized bool
}

var Framerate = 60.0
var Factory *factory = SetFactoryEngine(core.Impulse, when.Frequency(&Framerate))

func SetFactoryEngine(engine *core.Engine, potential core.Potential) *factory {
	f := &factory{}
	f.viewports = make(map[uint64]*Viewport)
	f.Dimension = temporal.DedicatedLoop(engine, potential, false, f.tick)
	return f
}

func (f *factory) NewViewport(title *string, size *std.XY[int], position *std.XY[int]) *Viewport {
	// Spark off a new open GL context on an impulse thread
	v := &Viewport{}
	v.ID = core.NextID()
	v.Size = size
	v.Title = title
	v.Position = position
	f.viewports[v.ID] = v
	return v
}

func (f *factory) HasViewports() bool {
	return len(f.viewports) > 0
}

func (f *factory) tick(ctx core.Context) {
	if len(f.viewports) == 0 {
		glfw.Terminate()
		f.initialized = false
		return
	}

	if !f.initialized {
		if err := glfw.Init(); err != nil {
			log.Fatalln("failed to initialize glfw:", err)
			return
		} else {
			f.initialized = true
		}
	}

	for _, vp := range f.viewports {
		var err error

		if vp.Window == nil {
			vp.Window, err = glfw.CreateWindow(vp.Size.X, vp.Size.Y, *vp.Title, nil, nil)
			vp.Window.SetPos(vp.Position.X, vp.Position.Y)
			if err != nil {
				panic(err)
			}
		}

		if vp.Window.ShouldClose() {
			vp.Destroyed = true
			delete(f.viewports, vp.ID)
			vp.Window.Destroy()
			continue
		}

		vp.Window.MakeContextCurrent()
		if err = gl.Init(); err != nil {
			log.Fatalln("failed to initialize gl:", err)
		}

		if !std.XYComparator(vp.lastSize, *vp.Size) {
			vp.Window.SetSize(vp.Size.X, vp.Size.Y)
			vp.lastSize = *vp.Size
		}

		if !std.XYComparator(vp.lastPosition, *vp.Position) {
			vp.Window.SetPos(vp.Position.X, vp.Position.Y)
			vp.lastPosition = *vp.Position
		}

		if !std.StringComparator(vp.lastTitle, *vp.Title) {
			vp.Window.SetTitle(*vp.Title)
			vp.lastTitle = *vp.Title
		}

		vp.Render(ctx)

		glfw.PollEvents()
		vp.Window.SwapBuffers()
	}
}
