package gateway

import (
	"github.com/gorilla/websocket"
	"github.com/jonavdm/cards-server/utils"
)

type Player struct {
	Name       string
	Connection *websocket.Conn
	Send       chan []byte
	Match      *Match
}

type Match struct {
	Players   []Player
	Join      chan string
	Broadcast chan []byte
}

type Hub struct {
	Matches map[string]*Match
	Create  chan Player
}

func (h *Hub) run() {
	for {
		p := <-h.Create

		c := utils.RandomString(6)
		m := &Match{
			Players:   make([]Player, 0),
			Join:      make(chan string),
			Broadcast: make(chan []byte),
		}

		p.Match = m
		m.Players = append(m.Players, p)

		h.Matches[c] = m

		p.Send <- []byte(c)
	}
}
