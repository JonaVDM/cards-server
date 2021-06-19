package gateway

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type App struct {
	Router   *mux.Router
	hub      *Hub
	upgrader websocket.Upgrader
}

func (a *App) Init() {
	a.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	a.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	a.hub = &Hub{
		Matches: make(map[string]*Match),
		Create:  make(chan Player),
	}
	go a.hub.run()

	a.Router.HandleFunc("/create", a.handleCreate())
	a.Router.HandleFunc("/join/{code}", a.handleJoin())
}

func (a *App) handleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := a.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		p := Player{
			Name:       "Bubble Head",
			Connection: conn,
			Send:       make(chan []byte),
			Match:      &Match{},
		}

		a.hub.Create <- p
	}
}

func (a *App) handleJoin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
