package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tinne26/kage-desk/display"
)

//go:embed shader.kage
var shaderProgram []byte

func main() {
	shader, err := ebiten.NewShader(shaderProgram)
	if err != nil {
		panic(err)
	}

	bounds := display.ImageSpiderCatDog().Bounds()

	game := &Game{shader: shader}

	ebiten.SetWindowTitle("Glitter")
	ebiten.SetWindowSize(bounds.Dx(), bounds.Dy()*2)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	err = ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}

type Game struct {
	shader   *ebiten.Shader
	vertices [4]ebiten.Vertex
}

func (g *Game) Layout(_, _ int) (int, int) {
	return 512, 512
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bounds := screen.Bounds()
	// top left:
	g.vertices[0].DstX = float32(bounds.Min.X)
	g.vertices[0].DstY = float32(bounds.Min.Y)

	// top right:
	g.vertices[1].DstX = float32(bounds.Max.X)
	g.vertices[1].DstY = float32(bounds.Min.Y)

	// bottom left:
	g.vertices[2].DstX = float32(bounds.Min.X)
	g.vertices[2].DstY = float32(bounds.Max.Y)

	// bottom right:
	g.vertices[3].DstX = float32(bounds.Max.X)
	g.vertices[3].DstY = float32(bounds.Max.Y)

	srcBounds := display.ImageSpiderCatDog().Bounds()
	// top left:
	g.vertices[0].SrcX = float32(srcBounds.Min.X)
	g.vertices[0].SrcY = float32(srcBounds.Min.Y)

	// top right:
	g.vertices[1].SrcX = float32(srcBounds.Max.X)
	g.vertices[1].SrcY = float32(srcBounds.Min.Y)

	// bottom left:
	g.vertices[2].SrcX = float32(srcBounds.Min.X)
	g.vertices[2].SrcY = float32(srcBounds.Max.Y)

	// bottom right:
	g.vertices[3].SrcX = float32(srcBounds.Max.X)
	g.vertices[3].SrcY = float32(srcBounds.Max.Y)

	var shaderOpts ebiten.DrawTrianglesShaderOptions
	shaderOpts.Images[0] = display.ImageSpiderCatDog()
	shaderOpts.Uniforms = make(map[string]interface{})
	shaderOpts.Uniforms["MirrorAlphaMult"] = float32(0.2)
	shaderOpts.Uniforms["VertDisplacement"] = 28

	indices := []uint16{0, 1, 2, 2, 1, 3}

	screen.DrawTrianglesShader(g.vertices[:], indices, g.shader, &shaderOpts)
}
