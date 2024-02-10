package balancer

import (
	"net/http"
	"net/url"

	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
)

type Balancer struct {
	URL      *url.URL
	Pool     *pool.ServerPool
	Strategy strategy.Strategy
	Done     chan int
}

func (b *Balancer) Handle(w http.ResponseWriter, r *http.Request) {
	b.Strategy.Handle(w, r, b.Pool)
}

func (b *Balancer) ListenAndServe() {
	http.ListenAndServe(b.URL.Host, http.HandlerFunc(b.Handle))
	b.Done <- 0
}
