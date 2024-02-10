package pool

type ServerPool struct {
	Backends []*Backend
}

func (sp *ServerPool) CheckHealth() {
	for _, b := range sp.Backends {
		b.CheckHealth()
	}
}
