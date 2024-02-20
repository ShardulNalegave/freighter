package strategy

import (
	"net/http"
	"sync/atomic"

	"github.com/ShardulNalegave/freighter/pool"
)

type RoundRobin struct {
	current uint64
}

func (rr *RoundRobin) Handle(r *http.Request, p *pool.ServerPool) *pool.Backend {
	next := rr.increment(p)
	l := len(p.Backends) + int(rr.current)
	for i := next; i < l; i++ {
		index := int(i) % len(p.Backends)
		if p.Backends[index].IsAlive() {
			if i != next {
				atomic.StoreUint64(&rr.current, uint64(index))
			}
			return p.Backends[index]
		}
	}

	return nil
}

func (rr *RoundRobin) increment(p *pool.ServerPool) int {
	return int(atomic.AddUint64(&rr.current, uint64(1)) % uint64(len(p.Backends)))
}
