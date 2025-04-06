package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/mouse"
	"github.com/ignite-laboratories/host/window"
	"time"
)

var scale = std.TimeScale[int]{Duration: time.Second * 2, Height: 5000}
var xCoords = temporal.Observer(core.Impulse, when.Frequency(std.HardRef(240.0).Ref), false, GetXCoords)
var yCoords = temporal.Observer(core.Impulse, when.Frequency(std.HardRef(240.0).Ref), false, GetYCoords)

func main() {
	viewport.NewWaveform(std.XY[int]{X: 640, Y: 480}, &scale, false, xCoords)
	viewport.NewWaveform(std.XY[int]{X: 640, Y: 480}, &scale, false, yCoords)
	core.Impulse.StopWhen(window.StopPotential)
	core.Impulse.Spark()
	core.ShutdownNow()
}

func GetXCoords[T int]() *T {
	x := T(mouse.Sample().Position.X)
	return &x
}

func GetYCoords[T int]() *T {
	y := T(mouse.Sample().Position.Y)
	return &y
}
