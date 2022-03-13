package endpoint

import "context"

type Endpoint[T any, K any] func(ctx context.Context, request T) (K, error)

type EndpointMiddleware[T any, K any] func(ctx context.Context, request T, next Endpoint[T, K]) (K, error)

type EndpointService[T any, K any] interface {
	Endpoint[T, K]
	EndpointMiddleware[T, K]
}
