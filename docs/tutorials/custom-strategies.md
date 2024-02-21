
# Custom Strategies
A Freighter strategy is any struct that implements the `Strategy` interface. This allows users to implement their own custom strategies which suit their needs. All they have to do is create a new struct which implements the interface and pass it to Freighter.

```go
type Strategy interface {
	Handle(r *http.Request, p *pool.ServerPool) *pool.Backend
}
```

## Re-implementing Round-Robin
Let's try to implement our own version of the Round-Robin strategy! Round-Robin is one of the simplest load-balancing strategies to distribute load across a set of backends. All it does is forward incoming requests to backends turn wise.

First, initialize a new project and add freighter as a dependency.
```bash
go mod init github.com/my/repo
go get github.com/ShardulNalegave/freighter
```

Now create `main.go` and add following code to it.
```go
package main

import (
	"net/url"
	"time"

	"github.com/ShardulNalegave/freighter"
	"github.com/ShardulNalegave/freighter/pool"
)

func main() {
	srv := freighter.NewFreighter(&freighter.Options{
		URL: &url.URL{
			Host: ":5000",
		},
		EnableConsoleLogging: true,
		HealthCheckInterval:  time.Second * 5,
		Backends: []*pool.Backend{
			pool.NewBackend(&url.URL{Host: ":8080", Scheme: "http"}, nil),
			pool.NewBackend(&url.URL{Host: ":8081", Scheme: "http"}, nil),
		},
    // TODO: Add our new strategy
	})

	srv.ListenAndServe()
}

```

Now, all thats let is to implement our strategy.
```go
type MyStrategy struct{
  current uint64
}
```

Taking a look at the `Strategy` interface, we see that `Handle` method should be defined on `MyStrategy`
```go
import (
	"net/http"
	"sync/atomic"

	"github.com/ShardulNalegave/freighter/pool"
)

func (s *MyStrategy) Handle(r *http.Request, p *pool.ServerPool) *pool.Backend {
	next := s.increment(p)
	l := len(p.Backends) + int(s.current)
	for i := next; i < l; i++ {
		index := int(i) % len(p.Backends)
		if p.Backends[index].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(index))
			}
			return p.Backends[index]
		}
	}

	return nil
}

func (s *MyStrategy) increment(p *pool.ServerPool) int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(p.Backends)))
}
```

What we are doing here is:-
1. When a request is received, get the next backend using the `current` index
2. Forward the request by using its reverse-proxy.
3. Atomically increment `current` by 1.
4. If `current` is equal to number to backends then reset `current` to 0.

<br />
Finally, pass our new strategy to the Freighter instance.

## Complete Code
```go
package main

import (
	"net/url"
	"time"

	"github.com/ShardulNalegave/freighter"
	"github.com/ShardulNalegave/freighter/pool"
)

func main() {
	srv := freighter.NewFreighter(&freighter.Options{
		URL: &url.URL{
			Host: ":5000",
		},
		EnableConsoleLogging: true,
		HealthCheckInterval:  time.Second * 5,
		Strategy:             &MyStrategy{}, // Our new strategy
		Backends: []*pool.Backend{
			pool.NewBackend(&url.URL{Host: ":8080", Scheme: "http"}, nil),
			pool.NewBackend(&url.URL{Host: ":8081", Scheme: "http"}, nil),
		},
	})

	srv.ListenAndServe()
}

type MyStrategy struct{
  current uint64
}

func (s *MyStrategy) Handle(r *http.Request, p *pool.ServerPool) *pool.Backend {
	next := s.increment(p)
	l := len(p.Backends) + int(s.current)
	for i := next; i < l; i++ {
		index := int(i) % len(p.Backends)
		if p.Backends[index].IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(index))
			}
			return p.Backends[index]
		}
	}

	return nil
}

func (s *MyStrategy) increment(p *pool.ServerPool) int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(p.Backends)))
}
```