
# Strategies
Load balancers have different strategies for distributing loads across a set of backends. Freighter provides built-in strategies but if needed you can always implement custom ones easily!

Following strategies are provided by Freighter:-
- [Round-Robin](/guide/strategies/round-robin)

## The `Strategy` interface
All strategies should implement the `Strategy` interface.

```go
type Strategy interface {
	Handle(r *http.Request, p *pool.ServerPool) *pool.Backend
}
```