package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/ignite-laboratories/support/ipsum"
	"github.com/ignite-laboratories/tiny"
)

var framerate = 60.0 //Hz

func main() {
	core.Verbose = true

	data := []byte(ipsum.Paragraph)
	viewport.NewBasicByteWave(core.Impulse, false, when.Frequency(&framerate), "Raw Bytes", nil, nil, data)

	phrase := tiny.NewPhrase(data...)
	fmt.Println(phrase.BitLength())
	phrase.QuarterSplit()
	fmt.Println(phrase.BitLength())

	bytes, _ := phrase.ToBytesAndBits()
	viewport.NewBasicByteWave(core.Impulse, false, when.Frequency(&framerate), "Quarter Split", nil, nil, bytes)

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
