package httpapi

import (
	"context"
	"encoding/json"
	"errors"
)

type CommandRouter struct {
	handlers map[string]CommandHandler
	decoders map[string]CommandDecoder
}

func NewCommandRouter() *CommandRouter {
	return &CommandRouter{
		handlers: make(map[string]CommandHandler),
		decoders: make(map[string]CommandDecoder),
	}
}

func (r *CommandRouter) Register(
	commandName string,
	handler CommandHandler,
	decoder CommandDecoder,
) {

	r.handlers[commandName] = handler
	r.decoders[commandName] = decoder
}

func (r *CommandRouter) Dispatch(
	ctx context.Context,
	commandName string,
	cmd any,
) error {

	handler, ok := r.handlers[commandName]
	if !ok {
		return errors.New("command handler not found")
	}

	return handler.Handle(ctx, cmd)
}

func (r *CommandRouter) ResolveCommandPayload(
	commandName string,
	data json.RawMessage,
) (any, error) {

	decoder, ok := r.decoders[commandName]
	if !ok {
		return nil, errors.New("decoder not found")
	}

	return decoder(data)
}
