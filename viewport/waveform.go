package viewport

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/host/hydra"
)

type Waveform[TValue any] struct {
	*hydra.WindowCtrl
}

func NewWaveform[TValue core.Numeric](title string, size std.XY[int], timeScale *std.TimeScale[TValue], isSigned bool, target *temporal.Dimension[TValue, any]) *Waveform[TValue] {
	wave := &Waveform[TValue]{}
	wave.WindowCtrl = hydra.CreateWindow(core.Impulse, title, size, std.XY[int]{X: 200, Y: 400}, wave.Render, when.Frequency(std.HardRef(60.0).Ref), false)

	return wave
}

func (w Waveform[TValue]) Render(ctx core.Context) {
	gl.ClearColor(0.2, 0.3, 0.4, 1.0) // RGB color
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (w Waveform[TValue]) Cleanup() {

}
