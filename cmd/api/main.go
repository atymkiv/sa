//
// Sa labs
//
// API Docs for sa labs
//     Schemes: http
//     Version: 0.0.1
//     Contact: Andriy Tymkiv <a.tymkiv99@gmail.com>
//     Host: localhost
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta

package main

import (
	"flag"
	gormsqlS "github.com/atymkiv/sa/cmd/api/shop/platform/gormsql"
	ss "github.com/atymkiv/sa/cmd/api/shop/service"
	st "github.com/atymkiv/sa/cmd/api/shop/transport"
	"github.com/atymkiv/sa/pkg/utl/config"
	"github.com/atymkiv/sa/pkg/utl/gorm"
	"github.com/atymkiv/sa/pkg/utl/messages"
	"github.com/atymkiv/sa/pkg/utl/nats"
	"github.com/atymkiv/sa/pkg/utl/server"
)

func main() {

	cfgPath := flag.String("p", "./config/api.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(Start(cfg))
}

func Start(cfg *config.Configuration) error {
	db, err := gorm.GetDbInstance(&cfg.DB)
	if err != nil {
		return err
	}

	e := server.New()
	natsClient, err := nats.New(cfg.Nats)
	messageService := messages.Create(natsClient)
	shopDB := gormsqlS.NewShop(db, )
	shopGr := e.Group("/shop")
	st.NewHTTP(ss.New(shopDB, messageService), shopGr)

	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})

	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
