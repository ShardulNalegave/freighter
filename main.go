package main

import (
	"encoding/json"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/ShardulNalegave/freighter/compose"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	confFileName := os.Args[1]
	conf, _ := os.ReadFile(confFileName)
	var c compose.ComposeConfig
	err := json.Unmarshal([]byte(conf), &c)
	if err != nil {
		log.Fatal().AnErr("Error", err)
	}

	backends := make([]*pool.Backend, 0)
	for _, rec := range c.Backends {
		u, _ := url.Parse(rec.Address)
		backends = append(backends, pool.NewBackend(u, rec.Metadata))
	}

	p := &pool.ServerPool{
		Backends: backends,
	}

	b := &balancer.Balancer{
		URL: &url.URL{
			Host: c.Address,
		},
		Pool:     p,
		Strategy: &strategy.RoundRobin{},
	}

	var wg sync.WaitGroup
	wg.Add(1)

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
