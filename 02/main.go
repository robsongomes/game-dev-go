package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.StrokeLine(screen, 200, 200, 400, 400, 3, color.White, true)
	vector.StrokeRect(screen, 450, 300, 150, 150, 3, color.White, true)
	vector.StrokeCircle(screen, 500, 200, 75, 3, color.White, true)
	vector.DrawFilledCircle(screen, 300, 200, 75, color.White, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello Ebitengine")
	game := Game{}
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
