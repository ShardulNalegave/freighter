package freighter

import (
	"time"

	"github.com/ShardulNalegave/freighter/pool"
)

func HealthCheck(p *pool.ServerPool, interval time.Duration) {
	p.CheckHealth()
	for range time.Tick(interval) {
		p.CheckHealth()
	}
}
