package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/hydra/sdl2"
	"github.com/ignite-laboratories/support/ipsum"
	"time"
)

var framerate = 60.0 //Hz

var view = viewport.NewBasicByteWave(core.Impulse, false, when.Frequency(&framerate), "Ipsum Byte Wave", nil, nil, []byte(ipsum.Paragraph[:256]))

func main() {
	core.Verbose = true
	data := ipsum.Generate(5)

	i := 0
	go func() {
		for core.Alive {
			ipsum := data[i : i+256]
			view.SetBytes([]byte(ipsum))

			i++
			if i > 1024 {
				i = 0
			}

			time.Sleep(time.Millisecond * 16)
		}
	}()

	core.Impulse.StopWhen(sdl2.HasNoWindows)
	core.Impulse.Spark()
}
