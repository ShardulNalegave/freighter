package pool

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}
