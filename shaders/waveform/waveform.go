package waveform

import (
	_ "embed"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/glitter"
	"log"
)

func Init() {
	fmt.Println("[glitter] initializing waveform shaders")
	if err := gl.Init(); err != nil {
		log.Fatalf("failed to initialize OpenGL: %v", err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	VertexShaderID = glitter.CompileShader(VertexShader, gl.VERTEX_SHADER)
	FragmentShaderID = glitter.CompileShader(FragmentShader, gl.FRAGMENT_SHADER)
	//GeometryShaderID = glitter.CompileShader(GeometryShader, gl.FRAGMENT_SHADER)
	SimpleProgram = glitter.LinkPrograms(VertexShaderID, FragmentShaderID)
	//GeometryProgram = glitter.LinkPrograms(VertexShaderID, GeometryShaderID, FragmentShaderID)
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
