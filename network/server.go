package network

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"

	"github.com/alexander-lazarov/lightriders/models"
)

var keyInput chan models.Direction
var boardOutput chan models.Board

var serverWriter *bufio.Writer

var serverEncoder *json.Encoder
var serverDecoder *json.Decoder

func InitServer() (*chan models.Direction, *chan models.Board) {
	fmt.Printf("Listening on port %d...\n", Port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", Port))
	if err != nil {
		panic("Network error!")
	}
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Network error! Trying again")
			continue
		}

		serverWriter = bufio.NewWriter(conn)
		serverEncoder = json.NewEncoder(serverWriter)
		serverDecoder = json.NewDecoder(bufio.NewReader(conn))

		keyInput = make(chan models.Direction)
		boardOutput = make(chan models.Board)

		keyReceiverInit()
		boardSenderInit()

		return &keyInput, &boardOutput
	}
}

func keyReceiverInit() {
	go func() {
		for {
			var keyToReceive models.Direction

			err := serverDecoder.Decode(&keyToReceive)
			if err != nil {
				panic("Network error - keyReceiverInit")
			}

			keyInput <- keyToReceive
		}
	}()
}

func boardSenderInit() {
	fmt.Printf("Init board sender\n")
	go func() {
		for {
			boardToSend := <-boardOutput

			err := serverEncoder.Encode(&boardToSend)

			if err != nil {
				panic("Network error - boardSenderInit")
			}

			serverWriter.Flush()
		}
	}()
}
