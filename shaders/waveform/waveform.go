package waveform

import (
	_ "embed"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/host/graphics"
	"log"
)

func Init() {
	fmt.Println("[glitter] - initializing waveform shaders")
	if err := gl.Init(); err != nil {
		log.Fatalf("Failed to initialize OpenGL: %v", err)
	}

	VertexShaderID = graphics.CompileShader(VertexShader, gl.VERTEX_SHADER)
	FragmentShaderID = graphics.CompileShader(FragmentShader, gl.FRAGMENT_SHADER)
	ProgramID = graphics.LinkPrograms(VertexShaderID, FragmentShaderID)
}

//go:embed waveform.vert
var VertexShader string
var VertexShaderID uint32

//go:embed waveform.frag
var FragmentShader string
var FragmentShaderID uint32

var ProgramID uint32
