package command

import (
	"context"
	"encoding/json"
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
