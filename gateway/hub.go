package gateway

import (
	"github.com/jonavdm/cards-server/utils"
)

type Hub struct {
	Matches map[string]*Match
	Create  chan Player
}

func (h *Hub) run() {
	for {
		p := <-h.Create

		c := utils.RandomString(3)
		m := &Match{
			Players: make([]Player, 0),
			Join:    make(chan Player),
			Leave:   make(chan Player),
		}

		go m.run()

		p.Match = m
		p.Send <- []byte(c)
		m.Join <- p
		h.Matches[c] = m
	}
}
