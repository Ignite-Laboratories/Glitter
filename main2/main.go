package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/hydra"
)

func init() {
	core.Verbose = true
}

var framerate = 60.0 //hz

func main() {
	var windowSize = &std.XY[int]{X: 320, Y: 240}

	viewport.NewTearing(true, when.Frequency(&framerate), "Mouse X", windowSize, nil)

	core.Impulse.StopWhen(hydra.HasNoWindows)
	core.Impulse.Spark()
}
