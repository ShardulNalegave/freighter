package balancer

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
	"github.com/rs/zerolog/log"
)

type Balancer struct {
	URL      *url.URL
	Pool     *pool.ServerPool
	Strategy strategy.Strategy
}

func (b *Balancer) Handle(w http.ResponseWriter, r *http.Request) {
	if srv := b.Strategy.Handle(r, b.Pool); srv != nil {
		srv.ReverseProxy.ServeHTTP(w, r)
	}
}

func (b *Balancer) ListenAndServe(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Info().Str("URL", b.URL.Host).Msg("Listening...")
	http.ListenAndServe(b.URL.Host, http.HandlerFunc(b.Handle))
}
