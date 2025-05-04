package main

import (
	"github.com/ignite-laboratories/core"
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
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))
	viewport.NewSDL2ScreenTearTester(sdl2.Create(core.Impulse, false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil))

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
