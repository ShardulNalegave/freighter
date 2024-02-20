package freighter

import (
	"time"

	"github.com/ShardulNalegave/freighter/pool"
)

func HealthCheck(p *pool.ServerPool, interval time.Duration) {
	t := time.NewTicker(interval)
	for range t.C {
		p.CheckHealth()
	}
}
