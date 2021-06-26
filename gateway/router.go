package gateway

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/jonavdm/cards-server/game"
)

type App struct {
	Router   *mux.Router
	hub      *game.Hub
	upgrader websocket.Upgrader
}

func (a *App) Init() {
	a.upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	a.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	a.hub = &game.Hub{
		Matches: make(map[string]*game.Match),
		Create:  make(chan *game.Player),
		Close:   make(chan string),
	}
	go a.hub.Run()

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

		p := game.Player{
			ID:         uuid.NewString(),
			Name:       "Bubble Head",
			Connection: conn,
			Send:       make(chan []byte),
			IsLeader:   true,
		}

		go p.Reader()
		go p.Writer()

		a.hub.Create <- &p
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

		p := game.Player{
			ID:         uuid.NewString(),
			Name:       "Pizza",
			Connection: conn,
			Send:       make(chan []byte),
			Match:      match,
			IsLeader:   false,
		}

		go p.Reader()
		go p.Writer()

		match.Join <- p
	}
}
