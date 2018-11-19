package conf

import (
	"github.com/caarlos0/env"
	"github.com/sirupsen/logrus"
)

type Log struct {
	Lvl string `env:"API_LOG_LEVEL" envDefault:"info"`
}

func (l *Log) GetLogEntry() *logrus.Entry {
	//err can be ignored in this case
	level, _ := logrus.ParseLevel(l.Lvl)

	logger := logrus.New()
	logger.SetLevel(level)

	return logrus.NewEntry(logger)
}

func (c *ConfigImpl) Log() *logrus.Entry {
	if c.log != nil {
		return c.log
	}

	c.Lock()
	defer c.Unlock()

	log := &Log{}
	if err := env.Parse(log); err != nil {
		panic(err)
	}

	c.log = log.GetLogEntry()

	return c.log
}
