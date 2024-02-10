package strategy

import (
	"net/http"
	"sync/atomic"

	"github.com/ShardulNalegave/freighter/pool"
)

type RoundRobin struct {
	current uint64
}

func (rr *RoundRobin) Handle(w http.ResponseWriter, r *http.Request, p *pool.ServerPool) {
	b := p.Backends[rr.current]
	b.ReverseProxy.ServeHTTP(w, r)

	atomic.AddUint64(&rr.current, uint64(1))
	if int(rr.current) == len(p.Backends) {
		atomic.StoreUint64(&rr.current, uint64(0))
	}
}
