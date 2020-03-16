package main

import (
	"fmt"
	"github.com/fwidjaya20/natsio/app-order/cmd/containers"
	"github.com/fwidjaya20/natsio/app-order/cmd/http"
	"github.com/fwidjaya20/natsio/app-order/config"
	"github.com/fwidjaya20/natsio/pkg/es"
	"github.com/go-chi/chi"
	stan "github.com/nats-io/go-nats-streaming"
	"github.com/oklog/oklog/pkg/group"
	http2 "net/http"
	"os"
)

func main() {
	var g group.Group

	httpContainer := containers.New()

	initHTTP(&g, httpContainer)
	initNATS()

	_ = g.Run()
}

func initHTTP(g *group.Group, container containers.Container) {
	HTTPAddress := config.GetEnv(config.HTTP_ADDR)

	var router *chi.Mux
	router = chi.NewRouter()

	router.Mount("/v1", http.MakeHandler(router, container))

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

func initNATS() {
	var natsAddr = config.GetEnv(config.NATS_ADDR)
	var natsClient = config.GetEnv(config.NATS_CLIENT)
	var natsCluster = config.GetEnv(config.NATS_CLUSTER)

	natsConn, err := stan.Connect(natsCluster, natsClient, stan.NatsURL(natsAddr))

	if nil != err {
		fmt.Println("transport", "nats", err)
		os.Exit(1)
	}

	es.SetGlobalPublisher(es.NewPublisher(natsConn))
}