package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/hydra"
	"github.com/ignite-laboratories/host/mouse"
	"time"
)

var framerate = 60.0 //hz
var xTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}
var yTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}

var xCoords = temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, SampleX)
var yCoords = temporal.Calculation(core.Impulse, when.Frequency(&mouse.SampleRate), false, SampleY)

func main() {
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	viewport.NewWaveform(false, when.Frequency(&framerate), "Mouse X", windowSize, nil, &xTimeScale, false, xCoords)
	viewport.NewWaveform(false, when.Frequency(&framerate), "Mouse Y", windowSize, nil, &yTimeScale, false, yCoords)

	core.Impulse.StopWhen(hydra.HasNoWindows)
	core.Impulse.Spark()
}

func SampleX(ctx core.Context) int {
	return mouse.Sample().Position.X
}

func SampleY(ctx core.Context) int {
	return mouse.Sample().Position.Y
}
