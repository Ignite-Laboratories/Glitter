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
	Size      *std.XY[int]
	Position  *std.XY[int]
	Title     *string
	Window    *glfw.Window
	Mutex     sync.Mutex
	Destroyed bool

	lastPosition std.XY[int]
	lastSize     std.XY[int]
	lastTitle    string
}

func New(title *string, size *std.XY[int], position *std.XY[int]) *Viewport {
	// Spark off a new open GL context on an impulse thread
	v := &Viewport{}
	v.Size = size
	v.Title = title
	v.Position = position
	return v
}

func (v *Viewport) Render(ctx core.Context) {

}
