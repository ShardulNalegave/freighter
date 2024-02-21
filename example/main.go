package main

import (
	"net/url"
	"time"

	"github.com/ShardulNalegave/freighter"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
)

func main() {
	srv := freighter.NewFreighter(&freighter.Options{
		URL: &url.URL{
			Host: ":5000",
		},
		EnableConsoleLogging: true,
		HealthCheckInterval:  time.Second * 5,
		Strategy:             &strategy.RoundRobin{},
		Backends: []*pool.Backend{
			pool.NewBackend(&url.URL{Host: ":8080", Scheme: "http"}, nil),
			pool.NewBackend(&url.URL{Host: ":8081", Scheme: "http"}, nil),
		},
	})

	srv.ListenAndServe()
}
