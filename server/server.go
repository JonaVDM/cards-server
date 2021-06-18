package server

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router.HandleFunc("/ping", a.handlePing())
}
