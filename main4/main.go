package main

import (
	"fmt"
	"github.com/ignite-laboratories/arwen"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
)

var framerate = 60.0 //Hz

func main() {
	core.Verbose = true

	source := []byte{50, 60, 70, 75, 70, 60, 50, 40, 35, 40}
	approximation := []byte{50, 75, 50, 40}
	delta := arwen.CreateDeltaWave(source, approximation)
	unsigned := arwen.UnsignDeltaWave(delta)
	fmt.Println(delta)
	fmt.Println(unsigned)

	bgColor, _ := std.RGBFromHex("44", "44", "44")
	view := viewport.NewStackedByteWave(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil, bgColor)

	view.AddBytes(std.RGBA{R: 1.0}, source)
	view.AddBytes(std.RGBA{G: 1.0}, approximation)
	view.AddBytes(std.RGBA{B: 1.0}, unsigned)

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
