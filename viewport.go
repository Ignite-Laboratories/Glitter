package glitter

import (
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"log"
	"sync"
)

func init() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
}

type Viewport struct {
	core.Entity
	Size      std.XY[int]
	Title     *string
	Window    *glfw.Window
	Mutex     sync.Mutex
	Destroyed bool

	invoke    []func()
	lastTitle string
}

func (v *Viewport) SetSize(size std.XY[int]) {
	v.callOnMainThread(func() {
		v.Window.SetSize(size.X, size.Y)
	})
}

func (v *Viewport) callOnMainThread(action func()) {
	v.Mutex.Lock()
	v.invoke = append(v.invoke, action)
	v.Mutex.Unlock()
}

func (v *Viewport) Render(ctx core.Context) {
}
