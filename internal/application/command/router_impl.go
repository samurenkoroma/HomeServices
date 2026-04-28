package command

import (
	"context"
	"encoding/json"
	"errors"
)

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

func (r *router) Register(cmd string, handler Handler, decoder DecoderFunc) {
	r.handlers[cmd] = handler
	r.decoders[cmd] = decoder
}

func (r *router) Dispatch(ctx context.Context, commandName string, cmd any) (any, error) {
	handler, ok := r.handlers[commandName]
	if !ok {
		return nil, errors.New("command handler not found")
	}
	return handler(ctx, cmd)
}

func (r *router) ResolveCommandPayload(commandName string, data json.RawMessage) (any, error) {
	decoder, ok := r.decoders[commandName]
	if !ok {
		return nil, errors.New("decoder not found")
	}

	return decoder(data)
}
