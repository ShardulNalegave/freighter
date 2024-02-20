package pool

// Server-Pool of all registered backends
type ServerPool struct {
	Backends []*Backend
}

// Checks health of all backends
func (sp *ServerPool) CheckHealth() {
	for _, b := range sp.Backends {
		b.CheckHealth()
	}
}
