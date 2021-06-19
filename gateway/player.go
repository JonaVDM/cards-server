package gateway

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	Name       string
	Connection *websocket.Conn
	Send       chan []byte
	Match      *Match
}

func (p *Player) Reader() {
	defer func() {
		p.Connection.Close()
	}()

	p.Connection.SetReadLimit(maxMessageSize)
	p.Connection.SetReadDeadline(time.Now().Add(pongWait))
	p.Connection.SetPongHandler(func(string) error {
		p.Connection.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, _, err := p.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
	}
}

func (p *Player) Writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		log.Print("Closing connection")
		ticker.Stop()
		p.Connection.Close()
	}()

	for {
		select {
		case msg := <-p.Send:
			p.Connection.SetWriteDeadline(time.Now().Add(writeWait))

			w, err := p.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(msg)

			l := len(p.Send)
			for i := 0; i < l; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-p.Send)
			}

			if err := w.Close(); err != nil {
				log.Print(err)
				return
			}

		case <-ticker.C:
			p.Connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := p.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
