// Package glitter provides a factory for creating graphical outputs.
package glitter

import (
	"github.com/ignite-laboratories/core"
)

var ModuleName = "glitter"

func init() {
	core.ModuleReport(ModuleName)
}

func Report() {}
