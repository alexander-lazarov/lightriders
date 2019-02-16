package game

import (
	"sync"
	"time"

	"github.com/alexander-lazarov/lightriders/graphics"
	"github.com/alexander-lazarov/lightriders/input"
	"github.com/alexander-lazarov/lightriders/models"
	"github.com/alexander-lazarov/lightriders/network"
	"github.com/hajimehoshi/ebiten"
)

var board models.Board
var boardLock sync.RWMutex

var gameType Type
var server string

var keyOutput *chan models.Direction
var boardInput *chan models.Board
var keyInput *chan models.Direction
var boardOutput *chan models.Board

// InitGame entry point of the game - inits the appriate game type
func InitGame(t Type, serverName string) {
	gameType = t

	switch gameType {
	case Server:
		initServerGame()
	case Client:
		initClientGame(serverName)
	case Hotseat:
		initHotseat()
	}
}

// initServerGame initializes server:
//   - inits network
//   - attaches input/outputs channels
//   - inits local game state
func initServerGame() {
	keyInput, boardOutput = network.InitServer()

	go func() {
		for {
			key := <-*keyInput

			boardLock.Lock()
			board.P2.SetDirection(key)
			boardLock.Unlock()
		}
	}()

	initLocalStateManager()
}

// initClientGame initializes client:
//   - inits network
//   - attaches input/output channels
func initClientGame(serverName string) {
	boardReady := make(chan struct{})

	keyOutput, boardInput = network.InitClient(serverName)

	board = *models.InitialBoard()

	graphics.InitBoardImage()

	go func() {
		receivedInput := false
		for {
			var boardReceived = <-*boardInput

			if !receivedInput {
				receivedInput = true
				boardReady <- struct{}{}
			}

			boardLock.Lock()
			board = boardReceived
			boardLock.Unlock()
		}
	}()

	<-boardReady

	graphics.InitBoardImage()

	ebiten.Run(update, graphics.TotalWidth, graphics.TotalHeight, 2, "Lightriders!")
}

// initLocalStateManager inits locally running loop
func initLocalStateManager() {
	board = *models.InitialBoard()

	graphics.InitBoardImage()
	go boardLoop()

	ebiten.Run(update, graphics.TotalWidth, graphics.TotalHeight, 2, "Lightriders!")
}

func initHotseat() {
	initLocalStateManager()
}

func handleKeyboardInput() {
	if gameType == Hotseat {
		input.HandleInputToPlayer(board.P1, &input.KeysetWASD)
		input.HandleInputToPlayer(board.P2, &input.KeysetArrows)
	} else if gameType == Server {
		input.HandleInputToPlayer(board.P1, &input.KeysetArrows)
	} else {
		key, ok := input.GetInputDir(&input.KeysetArrows)

		if ok {
			*keyOutput <- key
		}
	}
}

// update - game loop as from ebiten's prospective
// read input from the current frame, send/receive from network
// update screen state
func update(screen *ebiten.Image) error {
	boardLock.Lock()
	handleKeyboardInput()
	boardLock.Unlock()

	boardLock.RLock()
	graphics.UpdateBoardImage(&board)
	screen.DrawImage(graphics.GetBoardImage(), &ebiten.DrawImageOptions{})
	boardLock.RUnlock()

	return nil
}

// boardLoop holds the game logic - it advances the game
func boardLoop() {
	var winner models.Winner = models.NoWinner

	for {
		time.Sleep(speed)

		boardLock.Lock()
		if winner == models.NoWinner {
			winner, _ = board.Advance()
		}
		if gameType == Server {
			*boardOutput <- board
		}

		boardLock.Unlock()
	}
}
