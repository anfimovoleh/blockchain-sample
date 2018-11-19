package conf

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Config interface {
	HTTP() *HTTP
	Log() *logrus.Entry
	Block() *Block
}

type ConfigImpl struct {
	sync.Mutex

	//internal objects
	http  *HTTP
	log   *logrus.Entry
	block *Block
}

func New() Config {
	return &ConfigImpl{
		Mutex: sync.Mutex{},
	}
}
