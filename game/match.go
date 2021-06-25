package game

type Match struct {
	Players []Player
	Join    chan Player
	Leave   chan Player
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
			m.Players = append(m.Players, p)
			m.broadcast([]byte(p.Name + " Has joined the match"))

			if len(m.Players) == 4 {
				m.broadcast([]byte("4 Player are now in the match"))
			}
		case s := <-m.Leave:
			m.broadcast([]byte(s.Name + " Has disconnected"))
		}
	}
}
