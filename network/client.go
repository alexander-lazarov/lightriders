package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/alexander-lazarov/lightriders/models"
)

var keyOutput chan models.Direction
var boardInput chan models.Board

var clientReader *bufio.Reader
var clientWriter *bufio.Writer

var clientEncoder *json.Encoder
var clientDecoder *json.Decoder

// InitClient network and return two channels:
// - the first one consumes directions and sends them to the client
// - the second one consumes board states from the server
func InitClient(serverAddr string) (*chan models.Direction, *chan models.Board) {
	for {
		fmt.Printf("Connecting to %s:%d...", serverAddr, Port)
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverAddr, Port))
		if err != nil {
			fmt.Printf(" failed, retrying in 0.5s\n")
			time.Sleep(500 * time.Millisecond)
			continue
		}

		fmt.Println("Connected!")

		clientReader = bufio.NewReader(conn)
		clientWriter = bufio.NewWriter(conn)

		clientEncoder = json.NewEncoder(clientWriter)
		clientDecoder = json.NewDecoder(clientReader)

		keyOutput = make(chan models.Direction)
		boardInput = make(chan models.Board)

		keySenderInit()
		boardReceiverInit()

		return &keyOutput, &boardInput
	}
}

func keySenderInit() {
	go func() {
		for {
			keyToSend := <-keyOutput

			err := clientEncoder.Encode(&keyToSend)

			if err != nil {
				panic("Network error - keySenderInit")
			}

			clientWriter.Flush()
		}
	}()
}

func boardReceiverInit() {
	go func() {
		for {
			var receivedBoard models.Board

			err := clientDecoder.Decode(&receivedBoard)
			if err != nil {
				panic("Network error - boardReceiverInit")
			}

			boardInput <- receivedBoard
		}
	}()
}
