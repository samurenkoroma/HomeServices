package command

import (
	"context"
	"encoding/json"
	"errors"
)

type Handler interface {
	Handle(ctx context.Context, cmd any) error
	Name() string
}
type DecoderFunc func([]byte) (any, error)

type Router interface {
	Register(Handler, DecoderFunc)
	Dispatch(context.Context, string, any) error
	ResolveCommandPayload(string, json.RawMessage) (any, error)
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

func (r *router) Register(handler Handler, decoder DecoderFunc) {
	r.handlers[handler.Name()] = handler
	r.decoders[handler.Name()] = decoder
}

func (r *router) Dispatch(ctx context.Context, commandName string, cmd any) error {
	handler, ok := r.handlers[commandName]
	if !ok {
		return errors.New("command handler not found")
	}

	return handler.Handle(ctx, cmd)
}

func (r *router) ResolveCommandPayload(commandName string, data json.RawMessage) (any, error) {
	decoder, ok := r.decoders[commandName]
	if !ok {
		return nil, errors.New("decoder not found")
	}

	return decoder(data)
}
