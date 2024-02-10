package main

import (
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/ShardulNalegave/freighter/api"
	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	u1, _ := url.Parse("http://localhost:8080")
	p := &pool.ServerPool{
		Backends: []*pool.Backend{
			pool.NewBackend(u1, nil),
		},
	}

	b := &balancer.Balancer{
		URL: &url.URL{
			Host: ":5000",
		},
		Pool:     p,
		Strategy: &strategy.RoundRobin{},
	}

	a := api.NewAPI(&url.URL{
		Host: ":5001",
	}, b)

	var wg sync.WaitGroup
	wg.Add(2)

	go a.ListenAndServe(&wg)
	go b.ListenAndServe(&wg)
	go healthCheck(p)

	wg.Wait()
}

func healthCheck(p *pool.ServerPool) {
	t := time.NewTicker(5 * time.Second)
	for range t.C {
		log.Info().Msg("Running periodic Health-Check")
		p.CheckHealth()
	}
}
