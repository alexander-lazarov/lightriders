package graphics

import (
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/hajimehoshi/ebiten"
)

var boardImage *ebiten.Image
var cellImages [models.BoardWidth][models.BoardHeight](*ebiten.Image)

// InitBoardImage initializes the image that we will be drawing on
// and the initial cell images
func InitBoardImage() {
	img, _ := ebiten.NewImage(models.BoardWidth*tileWidth, models.BoardHeight*tileHeight, ebiten.FilterDefault)

	boardImage = img

	for i := range cellImages {
		for j := range cellImages[i] {
			tile, _ := ebiten.NewImage(tileWidth-2, tileHeight-2, ebiten.FilterDefault)

			cellImages[i][j] = tile
		}
	}
}

func GetBoardImage() *ebiten.Image {
	return boardImage
}

func UpdateBoardImage(board *models.Board) {
	for i := range cellImages {
		for j := range cellImages[i] {
			cellImages[i][j].Fill(getColor(board, models.Position{X: i, Y: j}))

			pos := &ebiten.GeoM{}
			pos.Translate((float64)(i*tileWidth), (float64)(j*tileHeight))

			boardImage.DrawImage(cellImages[i][j], &ebiten.DrawImageOptions{GeoM: *pos})
		}
	}
}
