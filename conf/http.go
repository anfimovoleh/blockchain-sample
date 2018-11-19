package conf

import (
	"fmt"

	"github.com/caarlos0/env"
)

type HTTP struct {
	Host string `env:"API_HTTP_HOST" envDefault:"localhost"`
	Port int    `env:"API_HTTP_PORT" envDefault:"8080"`
}

func (h *HTTP) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

func (c *ConfigImpl) HTTP() *HTTP {
	if c.http != nil {
		return c.http
	}

	c.Lock()
	defer c.Unlock()

	http := &HTTP{}
	if err := env.Parse(http); err != nil {
		panic(err)
	}

	c.http = http

	return c.http
}
