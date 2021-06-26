package game

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

type Player struct {
	ID         string
	Name       string
	Connection *websocket.Conn
	Send       chan []byte
	Match      *Match
	IsLeader   bool
}

func (p *Player) Reader() {
	defer func() {
		p.Match.Leave <- *p
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

			if _, err := w.Write(msg); err != nil {
				log.Print(err)
				return
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
