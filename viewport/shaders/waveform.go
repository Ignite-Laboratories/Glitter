package shaders

import (
	_ "embed"
	"fmt"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/ignite-laboratories/host/opengl"
)

func Init() {
	fmt.Println("[glitter] initializing waveform shaders")

	VertexShaderID = opengl.CompileShader(VertexShader, gl.VERTEX_SHADER)
	FragmentShaderID = opengl.CompileShader(FragmentShader, gl.FRAGMENT_SHADER)
	//GeometryShaderID = glitter.CompileShader(GeometryShader, gl.FRAGMENT_SHADER)
	SimpleProgram = opengl.LinkPrograms(VertexShaderID, FragmentShaderID)
	//GeometryProgram = glitter.LinkPrograms(VertexShaderID, GeometryShaderID, FragmentShaderID)

	gl.UseProgram(SimpleProgram)
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
