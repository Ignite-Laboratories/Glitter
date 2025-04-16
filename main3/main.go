package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/glfw"
)

func init() {
	core.Verbose = true
}

var framerate = 60.0 //hz

func main() {
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)
	viewport.NewScreenTearTesterGLFW(false, when.Frequency(&framerate), "Screen tearing test", windowSize, nil)

	core.Impulse.StopWhen(glfw.HasNoWindows)
	core.Impulse.Spark()
}
