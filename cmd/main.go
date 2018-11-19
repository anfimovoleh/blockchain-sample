package main

import (
	"time"

	"github.com/anfimovoleh/blockchain-sample"
	"github.com/anfimovoleh/blockchain-sample/core"
)

func main() {

	genesisBlock := core.Block{
		Index:     0,
		Timestamp: time.Now(),
	}

	core.Blockchain = append(core.Blockchain, genesisBlock)

	app.New().Start()
}
