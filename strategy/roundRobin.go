package strategy

import (
	"net/http"

	"github.com/ShardulNalegave/freighter/pool"
)

type RoundRobin struct {
	//
}

func (rr *RoundRobin) Handle(w http.ResponseWriter, r *http.Request, p *pool.ServerPool) {
	//
}
