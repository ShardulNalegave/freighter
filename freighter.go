package freighter

import (
	"net/url"
	"sync"
	"time"

	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
)

type Options struct {
	URL                 *url.URL
	Strategy            strategy.Strategy
	Backends            []*pool.Backend
	HealthCheckInterval time.Duration
}

type Freighter struct {
	p                   *pool.ServerPool
	b                   *balancer.Balancer
	healthCheckInterval time.Duration
}

func (f *Freighter) ListenAndServe() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f.b.ListenAndServe(&wg)
	go HealthCheck(f.p, f.healthCheckInterval)
	wg.Wait()
}

func NewFreighter(opts *Options) *Freighter {
	p := &pool.ServerPool{
		Backends: opts.Backends,
	}

	b := &balancer.Balancer{
		URL:      opts.URL,
		Pool:     p,
		Strategy: opts.Strategy,
	}

	return &Freighter{
		p:                   p,
		b:                   b,
		healthCheckInterval: opts.HealthCheckInterval,
	}
}
