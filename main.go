package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 640
)

type Game struct {
	offscreen    *ebiten.Image
	offscreenPix []byte
}

func NewGame() *Game {
	g := &Game{
		offscreen:    ebiten.NewImage(screenWidth, screenHeight),
		offscreenPix: make([]byte, screenWidth*screenHeight*4),
	}
	// Now it is not feasible to call updateOffscreen every frame due to performance.
	g.updateOffscreen(0.0, 0.0, 4)
	return g
}

func (gm *Game) updateOffscreen(centerX, centerY, size float64) {
	for j := 0; j < screenHeight; j++ {
		for i := 0; i < screenWidth; i++ {
			x := float64(i)*size/screenWidth - size/2 + centerX
			y := math.Sin(x)
			//screenY := (screenHeight-float64(y))*size/screenHeight - size/2 + centerY
			screenY := int((y - centerY + size/2) * screenHeight / size)

			r, g, b := 0x00, 0x00, 0x00
			if screenY == j {
				r, g, b = 0xff, 0x00, 0x00
			}

			// Check if the current pixel is on the x-axis
			if j == screenHeight/2 {
				r, g, b = 0x00, 0xff, 0x00 // Green color for the x-axis
			}

			// Check if the current pixel is on the y-axis
			if i == screenWidth/2 {
				r, g, b = 0x00, 0x00, 0xff // Blue color for the y-axis
			}

			// Handle the case where a pixel could belong to both axes and the sine wave
			if screenY == j && j == screenHeight/2 {
				r, g, b = 0xff, 0xff, 0x00 // Yellow color if on both sine wave and x-axis
			}
			if screenY == j && i == screenWidth/2 {
				r, g, b = 0xff, 0x00, 0xff // Magenta color if on both sine wave and y-axis
			}
			if j == screenHeight/2 && i == screenWidth/2 {
				r, g, b = 0x00, 0xff, 0xff // Cyan color if on both x-axis and y-axis
			}

			p := 4 * (i + j*screenWidth)
			gm.offscreenPix[p] = byte(r)
			gm.offscreenPix[p+1] = byte(g)
			gm.offscreenPix[p+2] = byte(b)
			gm.offscreenPix[p+3] = 0xff

		}
	}
	gm.offscreen.WritePixels(gm.offscreenPix)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.offscreen, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sin Graph")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
