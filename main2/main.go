package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/debugging"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
)

func init() {
	core.Verbose = true
}

var framerate = 60.0 //hz

func main() {
	core.Impulse.MaxFrequency = 1024.0
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	core.Verbosef("main thread", "%d\n", debugging.GetGoroutineID())

	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewTearing(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
