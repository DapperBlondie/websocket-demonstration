package main

import (
	"fmt"
	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.Handle("/", websocket.Handler(Echo))

	err := http.ListenAndServe("localhost:12345", nil)
	if err != nil {
		log.Log().Msg("Exited ...")
		os.Exit(1)
		return
	}
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
