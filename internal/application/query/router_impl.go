package query

import (
	"context"
	"sync"
)

// Decoder преобразует сырой payload (json, grpc, etc)

type router struct {
	mu       sync.RWMutex
	handlers map[string]registeredQuery
}

type registeredQuery struct {
	decoder Decoder
	handler HandlerFunc
}

func NewRouter() Router {
	return &router{
		handlers: make(map[string]registeredQuery),
	}
}

func (r *router) Register(name string, decoder Decoder, handler HandlerFunc) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.handlers[name] = registeredQuery{
		decoder: decoder,
		handler: handler,
	}
}

func (r *router) Dispatch(ctx context.Context, name string, payload []byte) (any, error) {

	r.mu.RLock()
	q, ok := r.handlers[name]
	r.mu.RUnlock()

	if !ok {
		return nil, ErrQueryNotFound
	}

	decoded, err := q.decoder(payload)
	if err != nil {
		return nil, err
	}

	return q.handler(ctx, decoded)
}
