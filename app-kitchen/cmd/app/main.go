package main

import (
	"fmt"
	"github.com/fwidjaya20/natsio/app-kitchen/cmd/nats"
	"github.com/fwidjaya20/natsio/app-kitchen/config"
	"github.com/fwidjaya20/natsio/pkg/es"
	"github.com/go-chi/chi"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/oklog/oklog/pkg/group"
	http2 "net/http"
	"os"
)

func main() {
	var g group.Group

	initHTTP(&g)
	initNATS(&g)

	_ = g.Run()
}

func initHTTP(g *group.Group) {
	HTTPAddress := config.GetEnv(config.HTTP_ADDR)

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/v1", nil)

	s := http2.Server{
		Addr: HTTPAddress,
		Handler: router,
	}

	g.Add(
		func() error {
			fmt.Println("transport", "debug/HTTP", "addr", HTTPAddress)
			return s.ListenAndServe()
		},
		func(err error) {
			if nil != err {
				fmt.Println("transport", "debug/HTTP", "addr", HTTPAddress, "error", err)
				panic(err)
			}
		},
	)
}

func initNATS(g *group.Group) {
	var natsAddr = config.GetEnv(config.NATS_ADDR)
	var natsClient = config.GetEnv(config.NATS_CLIENT)
	var natsCluster = config.GetEnv(config.NATS_CLUSTER)

	var subs []stan.Subscription

	natsConn, err := stan.Connect(natsCluster, natsClient, stan.NatsURL(natsAddr))

	if nil != err {
		fmt.Println("transport", "nats", err)
		os.Exit(1)
	}

	publisher := es.NewPublisher(natsConn)
	es.SetGlobalPublisher(publisher)

	g.Add(
		func() error {
			fmt.Println("transport", "nats", "addr", natsAddr)

			subs, err = nats.MakeSubscribers(natsConn)

			if nil != err {
				panic(err)
				return err
			}

			return nil
		},
		func(err error) {
			if nil != err {
				for _, sub := range subs {
					_ = sub.Close()
				}

				_ = natsConn.Close()
				panic(err)
			}
		},
	)
}