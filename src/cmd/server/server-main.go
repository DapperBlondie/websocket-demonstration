package main

import (
	"fmt"
	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))
	return
}

func Echo(ws *websocket.Conn) {
	fmt.Println("We are starting echoing ...")

	for j := 0; j < 10; j += 1 {
		msg := "Hello" + " " + strconv.Itoa(j)
		err := websocket.Message.Send(ws, msg)
		if err != nil {
			log.Error().Msg(err.Error() + "; Error in sending msg")
			return
		}
		var reply string
		err = websocket.Message.Receive(ws, &reply)
		if err != nil {
			log.Error().Msg(err.Error() + "; Error in sending msg")
			return
		}
		fmt.Println("Client said : " + reply)
	}
}
