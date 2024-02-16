package freighter

import (
	"net/url"
	"sync"

	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
)

type Config struct {
	URL      *url.URL
	Backends []*pool.Backend
	Strategy strategy.Strategy
}

type Freighter struct {
	p *pool.ServerPool
	b *balancer.Balancer
}

func (f *Freighter) ListenAndServe() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f.b.ListenAndServe(&wg)
	go HealthCheck(f.p)
	wg.Wait()
}

func NewFreighter(c *Config) *Freighter {
	p := &pool.ServerPool{
		Backends: c.Backends,
	}

	b := &balancer.Balancer{
		URL:      c.URL,
		Pool:     p,
		Strategy: c.Strategy,
	}

	return &Freighter{
		p: p,
		b: b,
	}
}
