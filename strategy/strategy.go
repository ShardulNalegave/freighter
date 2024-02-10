package strategy

import (
	"net/http"

	"github.com/ShardulNalegave/freighter/pool"
)

type Strategy interface {
	Handle(w http.ResponseWriter, r *http.Request, p *pool.ServerPool)
}
