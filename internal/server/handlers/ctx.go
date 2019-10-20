package handlers

import (
	"context"
	"github.com/anfimovoleh/blockchain-sample/core"
	"github.com/sirupsen/logrus"
	"net/http"
)

type CtxKey int

const (
	logCtxKey = iota
	ledgerCtxKey
)

func CtxLog(entry *logrus.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logrus.Entry {
	return r.Context().Value(logCtxKey).(*logrus.Entry)
}

func CtxLedger(ledger *core.Ledger) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, ledgerCtxKey, ledger)
	}
}

func Ledger(r *http.Request) *core.Ledger {
	return r.Context().Value(ledgerCtxKey).(*core.Ledger)
}
