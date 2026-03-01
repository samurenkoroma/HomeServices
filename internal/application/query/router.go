package query

import (
	"context"
	"errors"
)

type Handler func(ctx context.Context, query any) (any, error)
type DecoderFunc func(data []byte) (any, error)

type Router interface {
	Register(name string, handler Handler, decoder DecoderFunc)
	Dispatch(ctx context.Context, name string, data []byte) (any, error)
}
type router struct {
	handlers map[string]Handler
	decoders map[string]DecoderFunc
}

func NewRouter() Router {
	return &router{
		handlers: make(map[string]Handler),
		decoders: make(map[string]DecoderFunc),
	}
}

func (r *router) Register(
	name string,
	handler Handler,
	decoder DecoderFunc,
) {

	r.handlers[name] = handler
	r.decoders[name] = decoder
}
func (r *router) Dispatch(
	ctx context.Context,
	name string,
	data []byte,
) (any, error) {

	handler, ok := r.handlers[name]
	if !ok {
		return nil, errors.New("query handler not found")
	}

	decoder, ok := r.decoders[name]
	if !ok {
		return nil, errors.New("query decoder not found")
	}

	payload, err := decoder(data)
	if err != nil {
		return nil, err
	}

	return handler(ctx, payload)
}
