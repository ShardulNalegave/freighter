
# Backends
Backends are the servers for the load-balancer to balance.

## The `Backend` struct
Freighter options asks for an array of the provided `Backend` struct. This struct contains information about your backend. It also tracks its state and periodically pings to check its health after Freighter starts listening.

```go
// Backend is representation of an actual running backend.
type Backend struct {
	ID           string
	Metadata     interface{} // Metadata to be set by user
	URL          *url.URL
	alive        bool
	mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}
```

To create a instance of this struct, use the provided `NewBackend` function.

```go
b1 := pool.NewBackend(&url.URL{
  Host: ":8080",
  Scheme: "http",
}, nil)

b2 := pool.NewBackend(&url.URL{
  Host: ":8082",
  Scheme: "http",
}, nil)

b3 := pool.NewBackend(&url.URL{
  Host: ":8083",
  Scheme: "http",
}, nil)
```

## Metadata
It is possible to add metadata to these `Backend` instances. This is helpful as we often need to know more than just the URL of the backend while choosing a peer for an incoming request.

For example, we might want to use a strategy which forwards more requests to backends with higher processing power which makes a lot of sense. To do this, along with the setting the required strategy we will have to store more information about the backends as their metadata.

```go
type MyMetadata struct {
  MaxReqsInQueue int
}

b := pool.NewBackend(&url.URL{
  Host: ":8083",
  Scheme: "http",
}, MyMetadata{
  MaxReqsInQueue: 100
})
```