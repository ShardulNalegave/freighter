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

	u1, _ := url.Parse("http://localhost:8080")
	p := &pool.ServerPool{
		Backends: []*pool.Backend{
			pool.NewBackend(u1),
		},
	}

	b := &balancer.Balancer{
		URL: &url.URL{
			Host: ":5000",
		},
		Pool:     p,
		Done:     make(chan int),
		Strategy: &strategy.RoundRobin{},
	}

	log.Info().Str("URL", b.URL.Host).Msg("Listening...")
	go b.ListenAndServe()
	go healthCheck(p)

	<-b.Done
}

func healthCheck(p *pool.ServerPool) {
	t := time.NewTicker(5 * time.Second)
	for range t.C {
		log.Info().Msg("Running periodic Health-Check")
		p.CheckHealth()
	}
}
