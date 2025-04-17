package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/ignite-laboratories/support/ipsum"
	"github.com/ignite-laboratories/tiny"
)

var framerate = 60.0 //Hz

func main() {
	core.Verbose = true

	bgColor, _ := std.RGBFromHex("44", "44", "44")
	view := viewport.NewStackedByteWave(core.Impulse, false, when.Frequency(&framerate), "Stacked Byte Waves", nil, nil, bgColor)

	data := []byte(ipsum.Paragraph[:32])
	fgColorA, _ := std.RGBFromHex("FF", "A5", "5D")
	view.AddBytes(fgColorA, data)

	phrase := tiny.NewPhrase(data...)
	fmt.Println(phrase.BitLength())
	phrase.QuarterSplit()
	fmt.Println(phrase.BitLength())

	bytes, _ := phrase.ToBytesAndBits()
	fgColorB, _ := std.RGBFromHex("AC", "C5", "72")
	view.AddBytes(fgColorB, bytes)

	data2 := []byte(ipsum.Paragraph[32:48])
	fgColorC, _ := std.RGBFromHex("FF", "DF", "88")
	view.AddBytes(fgColorC, data2)

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
