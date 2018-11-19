package server

import (
	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/anfimovoleh/blockchain-sample/core"
	"github.com/anfimovoleh/blockchain-sample/internal/server/middlewares"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"

	"github.com/anfimovoleh/blockchain-sample/internal/server/handlers"
	"github.com/gorilla/mux"
)

func Router(cfg conf.Config, ledger *core.Ledger) http.Handler {
	r := mux.NewRouter()

	r.Use(
		middlewares.Logger(cfg.Log(), 15 * time.Second),
		middleware.Recoverer,
		middlewares.Ctx(
			handlers.CtxLog(cfg.Log()),
			handlers.CtxLedger(ledger),
		),
	)

	r.HandleFunc("/", handlers.GetBlockchain).Methods("GET")
	r.HandleFunc("/", handlers.AddRecords).Methods("POST")

	return r
}
