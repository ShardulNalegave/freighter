package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	u, _ := url.Parse("http://localhost:8080")
	rp := httputil.NewSingleHostReverseProxy(u)

	http.ListenAndServe(":5000", http.HandlerFunc(rp.ServeHTTP))
}
