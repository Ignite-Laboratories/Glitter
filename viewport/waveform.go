package viewport

import (
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/host/hydra"
	"time"
)

type Waveform[TValue any] struct {
	*hydra.WindowHead
	count int
}

func NewWaveform[TValue any, TField core.Numeric](framePotential core.Potential, title string, size std.XY[int], timeScale *std.TimeScale[TField], isSigned bool, target *temporal.Dimension[TValue, any], fieldSelector func(TValue) TField) *Waveform[TValue] {
	wave := &Waveform[TValue]{}
	wave.WindowHead = hydra.CreateWindow(core.Impulse, title, size, std.XY[int]{X: 200, Y: 400}, wave.Render, framePotential, false)

	go func() {
		for core.Alive && wave.Alive {
			wave.Window.SetTitle(fmt.Sprintf("%v - %d", title, wave.count))
			wave.count = 0
			time.Sleep(time.Second)
		}
	}()

	return wave
}

func (w *Waveform[TValue]) Render(ctx core.Context) {
	w.count++
	gl.ClearColor(0.2, 0.3, 0.4, 1.0) // RGB color
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (w *Waveform[TValue]) Cleanup() {

}
