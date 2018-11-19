package app

import (
	"net/http"
	"time"

	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/anfimovoleh/blockchain-sample/core"
	"github.com/anfimovoleh/blockchain-sample/internal/server"
	"github.com/pkg/errors"
)

type App struct {
	cfg    conf.Config
	ledger *core.Ledger
}

func New(config conf.Config) *App {
	return &App{
		cfg:    config,
		ledger: core.New(),
	}
}

func (a *App) Start() {
	server := http.Server{
		Addr:           a.cfg.HTTP().Addr(),
		Handler:        server.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(errors.Wrap(err, "failed to start app"))
	}
}
