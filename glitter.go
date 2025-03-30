package glitter

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"log"
)

// InitializeGLFW must be called before certain GLFW operations.
func InitializeGLFW() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
}

// InitializeGL must be called before certain GL operations.
func InitializeGL() {
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize gl:", err)
	}
}
