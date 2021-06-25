package gateway

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
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
		}

		go p.Reader()
		go p.Writer()

		a.hub.Create <- p
	}
}

func (a *App) handleJoin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		match, ok := a.hub.Matches[params["code"]]
		if !ok {
			w.Write([]byte("match not found"))
			return
		}

		conn, err := a.upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			w.Write([]byte("Error while upgrading"))
			return
		}

		p := Player{
			Name:       "Pizza",
			Connection: conn,
			Send:       make(chan []byte),
			Match:      match,
		}

		go p.Reader()
		go p.Writer()

		match.Join <- p
	}
}
