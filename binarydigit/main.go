package main

import (
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ignite-laboratories/glitter/assets"
)

func main() {
	shader, err := ebiten.NewShader(assets.Get.Shader("binarydigit/binarydigit.kage"))
	if err != nil {
		panic(err)
	}

	game := &Game{shader: shader}

	ebiten.SetWindowTitle("Glitter")
	ebiten.SetWindowSize(512, 512)
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

	var shaderOpts ebiten.DrawTrianglesShaderOptions
	indices := []uint16{0, 1, 2, 2, 1, 3}
	screen.DrawTrianglesShader(g.vertices[:], indices, g.shader, &shaderOpts)
}
