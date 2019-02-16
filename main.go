package main

import (
	"fmt"
	"os"

	"github.com/alexander-lazarov/lightriders/game"
)

func main() {
	gameType, server := game.GetType(os.Args)

	if gameType == game.Unknown {
		fmt.Println("Usage: ")
		fmt.Println("  lightriders [hotseat]       - 2 players on one computer")
		fmt.Println("  lightriders join            - host a multiplayer game")
		fmt.Println("  lightriders host <hostname> - join a multiplayer game")
	} else {
		game.InitGame(gameType, server)
	}
}
