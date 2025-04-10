package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/hydra"
	"github.com/ignite-laboratories/host/mouse"
	"time"
)

var framerate = 60.0 //hz
var xTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}
var yTimeScale = std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}

func main() {
	var windowSize = std.XY[int]{X: 320, Y: 240}
	viewport.NewWaveform(when.Frequency(&framerate), "1 Mouse X", windowSize, &xTimeScale, false, mouse.State, func(state std.MouseState) int { return state.Position.X })
	viewport.NewWaveform(when.Frequency(&framerate), "2 Mouse Y", windowSize, &yTimeScale, false, mouse.State, func(state std.MouseState) int { return state.Position.Y })

	core.Impulse.StopWhen(hydra.HasNoWindows)
	core.Impulse.Spark()
}
