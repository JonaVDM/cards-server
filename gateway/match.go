package gateway

import "log"

type Match struct {
	Players   []Player
	Join      chan Player
	Broadcast chan []byte
}

func (m *Match) broadCaster() {
	for {
		msg := <-m.Broadcast
		for _, p := range m.Players {
			log.Print("sending to ", p.Name)
			p.Send <- msg
		}
	}
}

func (m *Match) run() {
	for {
		p := <-m.Join

		m.Players = append(m.Players, p)
		// m.Broadcast <- []byte(p.Name + " Has joined the match")
		m.Players[0].Send <- []byte(p.Name + " has joined the match")

		if len(m.Players) == 4 {
			m.Broadcast <- []byte("4 Player are now in the match")
		}
	}
}
