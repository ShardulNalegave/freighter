package pool

import (
	"net"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
)

// Backend is representation of an actual running backend.
type Backend struct {
	ID           string
	Metadata     interface{} // Metadata to be set by user
	URL          *url.URL
	alive        bool
	mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

// Returns the alive status of backend
func (b *Backend) IsAlive() bool {
	b.mutex.RLock()
	alive := b.alive
	b.mutex.RUnlock()
	return alive
}

// Sets the alive status of backend
func (b *Backend) SetAlive(alive bool) {
	b.mutex.Lock()
	b.alive = alive
	b.mutex.Unlock()
}

// Checks backend health and update alive status accordingly
func (b *Backend) CheckHealth() bool {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", b.URL.Host, timeout)
	if err != nil {
		b.SetAlive(false)
		return false
	}

	_ = conn.Close()
	b.SetAlive(true)
	return true
}

// Constructs a new backend instance
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
