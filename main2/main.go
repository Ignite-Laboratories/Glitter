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
var position = std.XY[int]{X: 55, Y: 55}

func main() {
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)
	glitter.Factory.NewViewport(&title, &size, &position)

	core.Impulse.Loop(ChangeTitle, when.Frequency(std.HardRef(1.0).Ref), false)
	core.Impulse.StopWhen(func(ctx core.Context) bool {
		return !glitter.Factory.HasViewports()
	})
	core.Impulse.Spark()
}

var i = 0

func ChangeTitle(ctx core.Context) {
	title = fmt.Sprintf("Hello, World! %d", i)
	//position.X += i
	//position.Y += i
	i++
}
