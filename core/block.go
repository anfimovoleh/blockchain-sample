package core

import (
	"fmt"
	"time"

	"github.com/anfimovoleh/blockchain-sample/resources"
)

type Block struct {
	//Index current  number of block
	Index uint64

	//Timestamp is a time when block is applied to blockchain.
	//This value is automatically determined
	Timestamp time.Time

	//Record store data inside the block
	Records []resources.Record

	//Hash is a SHA256 identifier representing of current block
	Hash string

	//PrevHash is the SHA256 identifier of the previous block in the chain
	PrevHash string
}

func (b *Block) AddRecords(records ...resources.Record) {
	b.Records = append(b.Records, records...)
}

func (b Block) String() string {
	return fmt.Sprintf("%d%s%v%s",
		b.Index,
		b.Timestamp,
		b.Records,
		b.PrevHash,
	)
}
