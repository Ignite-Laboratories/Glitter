package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/mouse"
	"github.com/ignite-laboratories/hydra/glfw"
	"time"
)

var framerate = 60.0 //hz
var xTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}
var yTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}

var xCoords = temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, SampleX)
var yCoords = temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, SampleY)

func main() {
	core.Verbose = true
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	// TODO: Provide a better mouse system for GLFW

	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewBasicWaveformGLFW(core.Impulse, false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)

	core.Impulse.StopWhen(glfw.HasNoWindows)
	core.Impulse.Spark()
}

func SampleX(ctx core.Context) int {
	return mouse.Sample().Position.X
}

func SampleY(ctx core.Context) int {
	return mouse.Sample().Position.Y
}
