package gateway

import "github.com/gorilla/mux"

type App struct {
	Router *mux.Router
}

func (a *App) Init() {
	a.Router.HandleFunc("", a.handleConnection())
}
