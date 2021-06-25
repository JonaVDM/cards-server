package game

import (
	"github.com/jonavdm/cards-server/utils"
)

type Hub struct {
	Matches map[string]*Match
	Create  chan *Player
}

func (h *Hub) Run() {
	for {
		p := <-h.Create

		c := utils.RandomString(5)
		m := &Match{
			Players: make(map[string]Player),
			Join:    make(chan Player),
			Leave:   make(chan Player),
		}

		go m.run()

		p.Match = m
		p.Send <- []byte(c)
		m.Join <- *p
		h.Matches[c] = m
	}
}
