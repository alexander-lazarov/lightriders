package graphics

import (
	"image/color"

	"github.com/alexander-lazarov/lightriders/models"
)

func getColor(board *models.Board, x, y int) color.RGBA {
	cell := board.CellValue(x, y)

	switch cell {
	case models.Player1Head:
		return color.RGBA{R: 255, G: 0, B: 0, A: 255}
	case models.Player1Tail:
		return color.RGBA{R: 200, G: 0, B: 0, A: 255}
	case models.Player2Head:
		return color.RGBA{R: 0, G: 255, B: 0, A: 255}
	case models.Player2Tail:
		return color.RGBA{R: 0, G: 200, B: 0, A: 255}
	default:
		return color.RGBA{R: 0, G: 0, B: 0, A: 0}
	}
}
