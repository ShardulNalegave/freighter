package pool

import (
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Backend struct {
	ID           string
	Metadata     interface{}
	URL          *url.URL
	alive        bool
	mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

func (b *Backend) IsAlive() bool {
	b.mutex.RLock()
	alive := b.alive
	b.mutex.RUnlock()
	return alive
}

func (b *Backend) SetAlive(alive bool) {
	b.mutex.Lock()
	b.alive = alive
	b.mutex.Unlock()
}

func (b *Backend) CheckHealth() {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", b.URL.Host, timeout)
	if err != nil {
		log.Error().Str("URL", b.URL.Host).Msg("Backend not responding")
		b.SetAlive(false)
		return
	}

	_ = conn.Close()
	b.SetAlive(true)
}

func NewBackend(URL *url.URL, meta interface{}) *Backend {
	rp := httputil.NewSingleHostReverseProxy(URL)
	return &Backend{
		ID:           uuid.NewString(),
		Metadata:     meta,
		URL:          URL,
		ReverseProxy: rp,
		alive:        true,
		mutex:        sync.RWMutex{},
	}
}
