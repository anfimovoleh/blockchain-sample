package conf

import (
	"github.com/caarlos0/env"
)

type Block struct {
	Timeout     uint64 `env:"NODE_BLOCK_TIMEOUT" envDefault:"15"`
	MaxMsgCount uint64 `env:"NODE_BLOCK_MAX_MSG_COUNT" envDefault:"250"`
}

func (c *ConfigImpl) Block() *Block {
	if c.block != nil {
		return c.block
	}

	c.Lock()
	defer c.Unlock()

	block := &Block{}
	if err := env.Parse(block); err != nil {
		panic(err)
	}

	c.block = block

	return c.block
}
