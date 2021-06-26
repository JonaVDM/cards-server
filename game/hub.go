package game

import (
	"github.com/jonavdm/cards-server/utils"
)

type Hub struct {
	Matches map[string]*Match
	Create  chan *Player
	Close   chan string
}

func (h *Hub) Run() {
	for {
		select {
		case p := <-h.Create:

			c := utils.RandomString(5)
			m := &Match{
				Players: make(map[string]Player),
				Join:    make(chan Player),
				Leave:   make(chan Player),
				Hub:     h,
				Code:    c,
			}

			go m.run()

			p.Match = m
			p.Send <- []byte(c)
			m.Join <- *p
			h.Matches[c] = m

		case c := <-h.Close:
			delete(h.Matches, c)
		}
	}
}
