
# Freighter
Simple Load-Balancer written in Go-lang.

Freighter is not a pre-built binary/CLI-application, instead it is a Go library which you can add to your own binary.
This allows you to very easily extend its functionality and maybe even add new features which you require. Also, this helps avoid unnecessarily complicated configuration files.

**Documentation:** https://shardulnalegave.github.io/freighter

```go
package main

import (
	"net/url"
	"time"

	"github.com/ShardulNalegave/freighter"
	"github.com/ShardulNalegave/freighter/pool"
	"github.com/ShardulNalegave/freighter/strategy"
)

func main() {
	srv := freighter.NewFreighter(&freighter.Options{
		URL: &url.URL{
			Host: ":5000",
		},
		EnableConsoleLogging: true,
		HealthCheckInterval:  time.Second * 5,
		Strategy:             &strategy.RoundRobin{},
		Backends: []*pool.Backend{
			pool.NewBackend(&url.URL{Host: ":8080", Scheme: "http"}, nil),
			pool.NewBackend(&url.URL{Host: ":8081", Scheme: "http"}, nil),
		},
	})

	srv.ListenAndServe()
}
```