package glitter

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/glitter/viewport"
	"github.com/ignite-laboratories/host/graphics"
)

func init() {
	fmt.Println("[glitter]")
}

func NewWaveformWindow[TValue core.Numeric](target *temporal.Dimension[TValue, any]) *viewport.Waveform[TValue] {
	v := &viewport.Waveform[TValue]{}
	v.GLWindow = graphics.NewGLWindow(v)
	v.Dimension = target

	return v
}
