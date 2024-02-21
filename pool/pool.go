package pool

import "github.com/rs/zerolog"

// Server-Pool of all registered backends
type ServerPool struct {
	Backends []*Backend
	logger   *zerolog.Logger
}

func NewServerPool(b []*Backend, logger *zerolog.Logger) *ServerPool {
	return &ServerPool{
		Backends: b,
		logger:   logger,
	}
}

// Checks health of all backends
func (sp *ServerPool) CheckHealth() {
	for _, b := range sp.Backends {
		if !b.CheckHealth() {
			sp.logger.Warn().
				Str("Backend.ID", b.ID).
				Msg("Backend not responding")
		}
	}
}
