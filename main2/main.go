package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/glitter"
)

var title = "Hello, World!"
var size = std.XY[int]{X: 420, Y: 380}

func main() {
	glitter.CreateRed(title, size)
	glitter.CreateBlue(title, size)
	core.Impulse.StopWhen(glitter.DestroyedPotential)
	core.Impulse.Spark()
}
