package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter"
	"github.com/ignite-laboratories/host/mouse"
	"github.com/ignite-laboratories/host/window"
)

var xCoords = temporal.Observer(core.Impulse, when.Frequency(std.HardRef(120.0).Ref), false, GetXCoords)

func main() {
	glitter.NewWaveformWindow(xCoords)
	core.Impulse.StopWhen(window.StopPotential)
	core.Impulse.Spark()
}

func GetXCoords[T int]() *T {
	x := T(mouse.SampleCoordinates().X)
	return &x
}
