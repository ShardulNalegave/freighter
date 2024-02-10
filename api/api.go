package api

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/ShardulNalegave/freighter/balancer"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

type API struct {
	URL      *url.URL
	r        *chi.Mux
	Balancer *balancer.Balancer
}

func (a *API) initRoutes() {
	//
}

func (a *API) ListenAndServe(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Info().Str("URL", a.URL.Host).Msg("API Listening...")
	http.ListenAndServe(a.URL.Host, a.r)
}

func NewAPI(u *url.URL, b *balancer.Balancer) *API {
	a := &API{
		URL:      u,
		r:        chi.NewRouter(),
		Balancer: b,
	}
	a.initRoutes()
	return a
}
