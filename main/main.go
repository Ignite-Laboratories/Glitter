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
var freq = 240.0     //hz
var xCoords = temporal.Observer(core.Impulse, when.Frequency(&freq), false, GetXCoords)
var yCoords = temporal.Observer(core.Impulse, when.Frequency(&freq), false, GetYCoords)

func main() {
	var windowSize = std.XY[int]{X: 320, Y: 240}
	viewport.NewWaveform(&framerate, "1 Mouse X", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}).Ref, false, xCoords)
	viewport.NewWaveform(&framerate, "2 Mouse Y", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}).Ref, false, yCoords)
	viewport.NewWaveform(&framerate, "3 Mouse X", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}).Ref, false, xCoords)
	viewport.NewWaveform(&framerate, "4 Mouse Y", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}).Ref, false, yCoords)
	viewport.NewWaveform(&framerate, "5 Mouse X", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 2560}).Ref, false, xCoords)
	viewport.NewWaveform(&framerate, "6 Mouse Y", windowSize, std.HardRef(std.TimeScale[int]{Duration: time.Second * 2, Height: 1440}).Ref, false, yCoords)

	core.Impulse.StopWhen(hydra.When.HasNoWindows)
	core.Impulse.Spark()
}

func GetXCoords[T int]() *T {
	x := T(mouse.Sample().Position.X)
	return &x
}

func GetYCoords[T int]() *T {
	y := T(mouse.Sample().Position.Y)
	return &y
}
