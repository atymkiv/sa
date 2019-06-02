package nats

import (
	"github.com/atymkiv/sa/pkg/utl/config"
	"github.com/nats-io/go-nats"
)

func New(cfg *config.Nats) (*nats.Conn, error) {
	natsClient, err := nats.Connect(cfg.Host)

	return natsClient, err
}
