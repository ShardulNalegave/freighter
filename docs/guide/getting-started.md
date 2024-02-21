
# Getting Started
Freighter is not a pre-built binary/CLI-application, instead it is a Go library which you can add to your own binary.
This allows you to very easily extend its functionality and maybe even add new features which you require. Also, this helps avoid unnecessarily complicated configuration files.

## Installation
Firstly, initialize a Go module in the current directory.
```bash
go mod init github.com/my/repo
```

Then add freighter as a dependency.
```bash
go get github.com/ShardulNalegave/freighter
```

## Creating a Freighter instance
The `Freighter` struct handles almost all of the required work. You can create a new instance using the provided `NewFreighter` method and pass your options.

```go
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
```

## Start listening
Finally, start listening for requests.

```go
srv.ListenAndServe()
```