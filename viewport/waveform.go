package viewport

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/temporal"
	"github.com/ignite-laboratories/host/graphics"
)

type Waveform[TValue core.Numeric] struct {
	*graphics.GLWindow
	*temporal.Dimension[TValue, any]
}

func (v *Waveform[TValue]) Render() {
	v.Mutex.Lock()
	data := make([]std.Data[TValue], len(v.Timeline))
	copy(data, v.Timeline)
	v.Mutex.Unlock()

	// Clear the window with a background color
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}
