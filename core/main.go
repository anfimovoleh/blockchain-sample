package core

import (
	"sync/atomic"
	"time"

	"github.com/anfimovoleh/blockchain-sample/conf"
	"github.com/sirupsen/logrus"
)

type Ledger struct {
	conf.Block

	height uint64

	blocks []Block

	currentBlock Block

	//log object, needed to log ledger events
	log *logrus.Entry
}

func New(log *logrus.Entry, block *conf.Block) *Ledger {
	return &Ledger{
		height:       0,
		blocks:       []Block{},
		Block:        *block,
		currentBlock: Block{},
		log:          log,
	}
}

func (l *Ledger) Init(){
	genesisBlock := Block{
		Index:     0,
		Timestamp: time.Now(),
	}

	genesisBlock.Hash = CalculateHash(genesisBlock)

	l.blocks = append(l.blocks, genesisBlock)
}

func (l *Ledger) CloseBlock() {
	ticker := time.NewTicker(time.Duration(l.Timeout))

	for ; ; <-ticker.C {
		//Don't write block with empty records
		if len(l.currentBlock.Records) == 0 {
			continue
		}

		l.currentBlock.Index = l.LastBlock().Index + 1
		l.currentBlock.Timestamp = time.Now()
		l.currentBlock.PrevHash = l.LastBlock().Hash
		l.currentBlock.Hash = CalculateHash(l.currentBlock)

		//increment in thread safe manner
		atomic.AddUint64(&l.height, 1)

		//append current block to blockchain
		l.blocks = append(l.blocks, l.currentBlock)
		//resetting the current block
		l.currentBlock = Block{}
	}
}

func (l Ledger) Blocks() []Block {
	return l.blocks
}

func (l Ledger) LastBlock() Block {
	return l.blocks[l.height]
}

func (l *Ledger) CurrentBlock() *Block{
	return &l.currentBlock
}
