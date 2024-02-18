package strategy

import (
	"net/http"

	"github.com/ShardulNalegave/freighter/pool"
)

type Strategy interface {
	Handle(r *http.Request, p *pool.ServerPool) *pool.Backend
}
