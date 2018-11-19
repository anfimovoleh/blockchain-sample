package core

import (
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

func New(cfg conf.Config) *Ledger {
	return &Ledger{
		height:       0,
		blocks:       []Block{},
		Block:        *cfg.Block(),
		currentBlock: Block{},
		log:          cfg.Log(),
	}
}

func (l *Ledger) CloseBlock() error {
	ticker := time.NewTicker(l.Timeout)

	for ; ; <-ticker.C {
		//TODO FINISH ME
		break
	}
}

func (l *Ledger) Validate() error {
	//Don't write block with empty records
	if len(l.currentBlock.Records) == 0 {
		return ErrEmptyBlock
	}

	lastBlock := l.LastBlock()

	if lastBlock.Index+1 != l.currentBlock.Index {
		return ErrNotValidIndex
	}

	if lastBlock.Hash != l.currentBlock.PrevHash {
		return ErrNotValidPrevHash
	}

	if CalculateHash(lastBlock) != l.currentBlock.Hash {
		return ErrNotValidHash
	}
}

//func (l *Ledger) GenerateBlock() Block {
//
//	//take last block
//	oldBlock := l.blocks[l.height]
//
//	newBlock := Block{
//		Index:     oldBlock.Index + 1,
//		Timestamp: time.Now(),
//		Records:   record,
//		PrevHash:  oldBlock.Hash,
//	}
//
//	newBlock.Hash = CalculateHash(newBlock)
//
//	return newBlock
//}

func (l Ledger) LastBlock() Block {
	return l.blocks[l.height]
}
