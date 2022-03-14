package endpoint

import "context"

type Endpoint[T any, K any] func(ctx context.Context, request T) (K, error)

type EndpointMiddleware[T any, K any] func(ctx context.Context, request T, next Endpoint[T, K]) (K, error)

func ChainEndpointMiddleware[T any, K any](endpoint Endpoint[T, K], middlewares ...EndpointMiddleware[T, K]) Endpoint[T, K] {
	if len(middlewares) == 0 {
		return endpoint
	}
	return func(ctx context.Context, request T) (K, error) {
		return middlewares[0](ctx, request, ChainEndpointMiddleware(endpoint, middlewares[1:]...))
	}
}
