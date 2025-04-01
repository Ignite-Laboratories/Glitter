package glitter

import (
	"fmt"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"log"
	"runtime"
	"sync"
)

type Viewport struct {
	core.Entity
	Size        std.XY[int]
	Title       string
	Mutex       sync.Mutex
	Destroyed   bool
	initialized bool

	window  *sdl.Window
	impulse chan std.ChannelAction
}

func CreateViewport(title string, size std.XY[int], renderAction core.Action) *Viewport {
	var err error

	v := &Viewport{}
	v.ID = core.NextID()
	v.Size = size
	v.Title = title
	v.impulse = make(chan std.ChannelAction)
	Viewports[v.ID] = v

	if core.Verbose {
		fmt.Printf("Sparking Viewport - [%d] %v\n", v.ID, v.Title)
	}
	go func() {
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		for !initialized {
		}

		var wg sync.WaitGroup
		wg.Add(1)
		*Dimension.Cache <- std.ChannelAction{Action: func() {
			v.window, err = sdl.CreateWindow(v.Title, 640, 480, sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE)
			if err != nil {
				panic(err)
			}
			wg.Done()
		}}
		defer v.window.Destroy()
		wg.Wait()
		v.initialized = true

		if err = gl.Init(); err != nil {
			log.Fatalf("Failed to initialize OpenGL bindings: %v", err)
		}

		var context sdl.GLContext
		context, err = sdl.GL_CreateContext(v.window)
		if err != nil {
			panic(err)
		}
		defer sdl.GL_DestroyContext(context)

		for msg := range v.impulse {
			if v.Destroyed {
				break
			}

			if msg.Action != nil {
				msg.Action()
				continue
			}

			if !std.StringComparator(v.Title, v.window.Title()) {
				v.window.SetTitle(v.Title)
			}

			if err = sdl.GL_MakeCurrent(v.window, context); err != nil {
				panic(err)
			}

			renderAction(msg.Context)
			msg.WaitGroup.Done()
		}

		v.window.Destroy()
		delete(Viewports, v.ID)
	}()
	return v
}
