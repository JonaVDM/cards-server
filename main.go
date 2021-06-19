package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonavdm/cards-server/gateway"
	"github.com/jonavdm/cards-server/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router := mux.NewRouter().StrictSlash(false)

	srv := server.App{
		Router: router.PathPrefix("/api").Subrouter(),
	}
	srv.Init()

	gw := gateway.App{
		Router: router.PathPrefix("/gw").Subrouter(),
	}
	gw.Init()

	s := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
