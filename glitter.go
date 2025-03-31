package glitter

import (
	"fmt"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"log"
)

func init() {
	fmt.Println("init - [glitter]")
}

// Framerate represents the global framerate for all rendering.
var Framerate = 240.0

// Dimension represents the underlying dimension that drives all rendering.
var Dimension = temporal.DedicatedLoop(core.Impulse, when.Frequency(&Framerate), false, tick)

// Viewports holds the currently active viewports - when they are destroyed, they are removed from this map.
var Viewports map[uint64]*Viewport = make(map[uint64]*Viewport)

// Spark creates a new basic viewport window that renders nothing.
func Spark(title *string, size std.XY[int]) *Viewport {
	// Spark off a new open GL context on an impulse thread
	v := &Viewport{}
	v.ID = core.NextID()
	v.Size = size
	v.Title = title
	v.invoke = make([]func(), 0)
	fmt.Printf("Sparking Viewport - [%d] %v\n", v.ID, *v.Title)
	Viewports[v.ID] = v
	return v
}

var glInitialized bool

func tick(ctx core.Context) {
	if len(Viewports) == 0 {
		glfw.Terminate()
		glInitialized = false
		return
	}

	if !glInitialized {
		if err := glfw.Init(); err != nil {
			log.Fatalln("failed to initialize glfw:", err)
			return
		} else {
			glInitialized = true
		}
	}

	for _, v := range Viewports {
		var err error

		if v.Window == nil {
			v.Window, err = glfw.CreateWindow(v.Size.X, v.Size.Y, *v.Title, nil, nil)
			if err != nil {
				panic(err)
			}
		}

		if v.Window.ShouldClose() {
			v.Destroyed = true
			fmt.Printf("Destroying Viewport - [%d] %v\n", v.ID, *v.Title)
			delete(Viewports, v.ID)
			v.Window.Destroy()
			continue
		}

		if !std.StringComparator(v.lastTitle, *v.Title) {
			v.Window.SetTitle(*v.Title)
			v.lastTitle = *v.Title
		}

		// Set the context and initialize GL
		v.Window.MakeContextCurrent()
		if err = gl.Init(); err != nil {
			log.Fatalln("failed to initialize gl:", err)
		}

		v.Mutex.Lock()
		for _, act := range v.invoke {
			act()
		}
		v.invoke = make([]func(), 0)
		v.Mutex.Unlock()

		v.Render(ctx)

		glfw.PollEvents()
		v.Window.SwapBuffers()
	}
}
