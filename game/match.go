package game

import (
	"fmt"
	"log"
)

type Match struct {
	Players map[string]Player
	Join    chan Player
	Leave   chan Player
}

func (m *Match) leaderCount() int {
	counter := 0

	for _, p := range m.Players {
		if p.IsLeader {
			counter++
		}
	}

	return counter
}

func (m *Match) broadcast(msg []byte) {
	for _, p := range m.Players {
		p.Send <- msg
	}
}

func (m *Match) run() {
	for {
		select {
		case p := <-m.Join:
			m.Players[p.ID] = p
			m.broadcast([]byte(p.Name + " Has joined the match"))

			if len(m.Players) == 4 {
				m.broadcast([]byte("4 Player are now in the match"))
			}
		case p := <-m.Leave:
			delete(m.Players, p.ID)
			m.broadcast([]byte(p.Name + " Has disconnected"))

			if len(m.Players) == 0 {
				log.Print("No players left, closing match (TODO)")

				return
			}

			if m.leaderCount() == 0 {
				for _, p := range m.Players {
					p.IsLeader = true
					m.broadcast([]byte(
						fmt.Sprintf("%s (%s) is now the leader", p.Name, p.ID),
					))
					break
				}
			}
		}
	}
}
