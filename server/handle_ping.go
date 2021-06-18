package server

import (
	"net/http"
)

func (a *App) handlePing() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Pong!"))
	}
}
