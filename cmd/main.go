package main

import (
	"github.com/anfimovoleh/blockchain-sample"
	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/pkg/errors"
)

func main() {
	cfg := conf.New()

	app := app.New(cfg)

	if err := app.Start(); err != nil{
		panic(errors.Wrap(err, "failed to start app"))
	}
}
