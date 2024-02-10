package main

import (
	"net/url"
	"os"
	"time"

	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	p := &pool.ServerPool{
		Backends: make([]*pool.Backend, 0),
	}

	b := &balancer.Balancer{
		URL: &url.URL{
			Host: ":5000",
		},
		Pool:     p,
		Done:     make(chan bool),
		Strategy: &strategy.RoundRobin{},
	}

	b.ListenAndServe()
	go healthCheck(p)

	<-b.Done
}

func healthCheck(p *pool.ServerPool) {
	t := time.NewTicker(20 * time.Second)
	for range t.C {
		p.CheckHealth()
	}
}
