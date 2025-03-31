package main

import (
	"fmt"
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter"
)

var title = "Hello, World!"
var size = std.XY[int]{X: 420, Y: 380}

var vp = glitter.Spark(&title, size)

func main() {
	glitter.Spark(&title, size)
	glitter.Spark(&title, size)
	glitter.Spark(&title, size)
	glitter.Spark(&title, size)
	glitter.Spark(&title, size)
	glitter.Spark(&title, size)

	core.Impulse.Loop(ChangeTitle, when.Frequency(std.HardRef(1.0).Ref), false)
	core.Impulse.StopWhen(func(ctx core.Context) bool {
		return len(glitter.Viewports) == 0
	})
	core.Impulse.Spark()
}

var i = 0

var toggle bool

func ChangeTitle(ctx core.Context) {
	title = fmt.Sprintf("Hello, World! %d", i)
	if toggle {
		toggle = false
		vp.SetSize(std.XY[int]{X: 320, Y: 240})
	} else {
		toggle = true
		vp.SetSize(std.XY[int]{X: 640, Y: 480})
	}
	i++
}
