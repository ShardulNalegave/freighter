package freighter

import (
	"net/http"
	"net/url"
	"sync"
	"time"

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
	URL                 *url.URL
	pl                  *pool.ServerPool
	Strategy            strategy.Strategy
	healthCheckInterval time.Duration
}

func (f *Freighter) ListenAndServe(wg *sync.WaitGroup) {
	defer wg.Done()
	go HealthCheck(f.pl, f.healthCheckInterval)
	http.ListenAndServe(f.URL.Host, http.HandlerFunc(f.Handle))
}

func (f *Freighter) Handle(w http.ResponseWriter, r *http.Request) {
	if backend := f.Strategy.Handle(r, f.pl); backend != nil {
		backend.ReverseProxy.ServeHTTP(w, r)
	} else {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
	}
}

func NewFreighter(opts *Options) *Freighter {
	pl := &pool.ServerPool{
		Backends: opts.Backends,
	}

	return &Freighter{
		URL:                 opts.URL,
		Strategy:            opts.Strategy,
		pl:                  pl,
		healthCheckInterval: opts.HealthCheckInterval,
	}
}
