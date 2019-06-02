package main

import (
	"github.com/atymkiv/sa/cmd/subscribers/service"
	"github.com/atymkiv/sa/pkg/utl/config"
	"github.com/atymkiv/sa/pkg/utl/nats"
	"sync"
)

func main() {
	cfg, err := config.Load("./config/subscribers.yaml")
	checkErr(err)

	natsClient, err := nats.New(cfg.Nats)
	checkErr(err)

	service.Subscriber(natsClient, "shops")

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
