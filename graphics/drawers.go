package graphics

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

// DrawBoard creates a new image and draws the current state
func DrawBoard(board *models.Board) *ebiten.Image {
	img, _ := ebiten.NewImage(models.BoardWidth*tileWidth, models.BoardHeight*tileHeight, ebiten.FilterDefault)

	for i := 0; i < models.BoardWidth; i++ {
		for j := 0; j < models.BoardWidth; j++ {
			tile, _ := ebiten.NewImage(tileWidth-2, tileHeight-2, ebiten.FilterDefault)
			tile.Fill(getColor(board, models.Position{X: i, Y: j}))

			pos := &ebiten.GeoM{}
			pos.Translate((float64)(i*tileWidth), (float64)(j*tileHeight))

			img.DrawImage(tile, &ebiten.DrawImageOptions{GeoM: *pos})
		}
	}

	return img
}
