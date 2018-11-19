package server

import (
	"net/http"

	"github.com/anfimovoleh/blockchain-sample/internal/server/handlers"
	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.GetBlockchain).Methods("GET")
	router.HandleFunc("/", handlers.WriteBlock).Methods("POST")

	return router
}
