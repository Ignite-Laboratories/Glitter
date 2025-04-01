package glitter

import (
	"fmt"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("init - [glitter]")
	go spark()
}

// Framerate represents the global framerate for all rendering.
var Framerate = 10.0

var Potential = when.Frequency(&Framerate)

// Dimension represents the underlying dimension that drives all rendering.
var Dimension = temporal.ChannelLoop(core.Impulse, func(ctx core.Context) bool {
	return Potential(ctx)
}, false)

// Viewports holds the currently active viewports - when they are destroyed, they are removed from this map.
var Viewports map[uint64]*Viewport = make(map[uint64]*Viewport)

var initialized bool
var destroy bool

func DestroyedPotential(ctx core.Context) bool {
	return destroy
}

func spark() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	defer binsdl.Load().Unload()
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	initialized = true

	for msg := range *Dimension.Cache {
		if destroy {
			break
		}
		if msg.Action != nil {
			msg.Action()
			continue
		}
		tick(msg.Context)
		if destroy {
			break
		}
	}

	for _, v := range Viewports {
		if v.window != nil {
			v.Destroyed = true
		}
	}
}

func tick(ctx core.Context) {
	var event sdl.Event
	for sdl.PollEvent(&event) {
		switch event.Type {
		case sdl.EVENT_WINDOW_CLOSE_REQUESTED:
			evt := event.WindowEvent()
			closeCount := 0
			for _, v := range Viewports {
				id, _ := v.window.ID()
				if id == evt.WindowID {
					v.Destroyed = true
					v.impulse <- std.ChannelAction{Action: func() {}}
					closeCount++
				}
			}
			if closeCount == len(Viewports) {
				destroy = true
			}
			return
		}
	}
	var wg sync.WaitGroup
	for _, v := range Viewports {
		if v.initialized && !v.Destroyed {
			wg.Add(1)
			v.impulse <- std.ChannelAction{Context: ctx, WaitGroup: &wg}
		}
	}

	wg.Wait()
}
