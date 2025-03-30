package main

import (
	"github.com/ignite-laboratories/core"
	"github.com/ignite-laboratories/core/std"
	"github.com/ignite-laboratories/core/when"
	"github.com/ignite-laboratories/glitter"
)

var viewport = glitter.NewViewport(core.Impulse, when.Frequency(std.HardRef(30.0).Ref), std.HardRef("Glitter").Ref, std.HardRef(std.XY[int]{X: 420, Y: 380}).Ref)

func main() {
	core.Impulse.StopWhen(std.PotentialTarget(&viewport.Destroyed))
	core.Impulse.Spark()
}
