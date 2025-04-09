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
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	VertexShaderID = graphics.CompileShader(VertexShader, gl.VERTEX_SHADER)
	FragmentShaderID = graphics.CompileShader(FragmentShader, gl.FRAGMENT_SHADER)
	//GeometryShaderID = graphics.CompileShader(GeometryShader, gl.FRAGMENT_SHADER)
	SimpleProgram = graphics.LinkPrograms(VertexShaderID, FragmentShaderID)
	//GeometryProgram = graphics.LinkPrograms(VertexShaderID, GeometryShaderID, FragmentShaderID)
}

//go:embed waveform.vert
var VertexShader string
var VertexShaderID uint32

//go:embed waveform.frag
var FragmentShader string
var FragmentShaderID uint32

//go:embed waveform.geom
var GeometryShader string
var GeometryShaderID uint32

var SimpleProgram uint32
var GeometryProgram uint32
