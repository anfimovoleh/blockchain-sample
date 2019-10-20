package app

import (
	"net/http"
	"time"

	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/anfimovoleh/blockchain-sample/core"
	"github.com/anfimovoleh/blockchain-sample/internal/server"
)

type App struct {
	cfg    conf.Config
	ledger *core.Ledger
}

func New(cfg conf.Config) *App {
	ledger := core.New(cfg.Log(), cfg.Block())
	ledger.Init()

	return &App{
		cfg:    cfg,
		ledger: ledger,
	}
}

func (a *App) Start() error {
	go a.ledger.CloseBlock()

	server := http.Server{
		Addr:           a.cfg.HTTP().Addr(),
		Handler:        server.Router(a.cfg, a.ledger),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
