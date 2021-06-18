package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jonavdm/cards-server/server"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	router := mux.NewRouter().StrictSlash(false)
	srvRouter := router.PathPrefix("/api").Subrouter()

	srv := server.App{
		Router: srvRouter,
	}
	srv.Init()

	s := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
