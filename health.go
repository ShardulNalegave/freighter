package freighter

import (
	"time"

	"github.com/ShardulNalegave/freighter/pool"
	"github.com/rs/zerolog/log"
)

func HealthCheck(p *pool.ServerPool) {
	t := time.NewTicker(5 * time.Second)
	for range t.C {
		log.Info().Msg("Running periodic Health-Check")
		p.CheckHealth()
	}
}
