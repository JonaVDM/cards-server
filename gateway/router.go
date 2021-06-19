package gateway

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router.HandleFunc("/new", a.handleConnection())
	a.Router.HandleFunc("/random", a.handleConnection())
	a.Router.HandleFunc("/join/{code}", a.handleConnection())
}

func (a *App) handleConnection() http.HandlerFunc {
	hub := newHub()
	go hub.run()

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return func(w http.ResponseWriter, r *http.Request) {
		// code := mux.Vars(r)

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{hub, conn, make(chan []byte, 256)}
		client.hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
