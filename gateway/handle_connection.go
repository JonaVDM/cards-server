package gateway

import (
	"fmt"
	"log"
	"net/http"
)

func (a *App) handleConnection() http.HandlerFunc {
	h := newHub()
	go h.run()

	fmt.Println("Started the ws")

	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("connection incomming")

		conn, err := upgrader.Upgrade(rw, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{h, conn, make(chan []byte, 256)}
		client.hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
