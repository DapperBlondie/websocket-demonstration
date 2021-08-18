package main

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
	"io"
	"os"
)

func main() {
	conn, err := websocket.Dial("ws://localhost:12345", "", "http://localhost:12345")
	if err != nil {
		log.Error().Msg(err.Error())
		return
	}
	var msg string
	for {
		err = websocket.Message.Receive(conn, &msg)
		if err != nil {
			if err == io.EOF {
				log.Log().Msg("")
				os.Exit(1)
			}
			log.Error().Msg(err.Error() + "; in receiving data")
			return
		}
		err = websocket.Message.Send(conn, msg)
		if err != nil {
			log.Log().Msg("Could not be able to return msg")
			break
		}
	}
	return
}
