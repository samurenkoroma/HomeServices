package command

import (
	"context"
	"encoding/json"
)

type Handler interface {
	Handle(ctx context.Context, cmd any) error
}
type DecoderFunc func([]byte) (any, error)

type Router interface {
	Register(string, Handler, DecoderFunc)
	Dispatch(context.Context, string, any) error
	ResolveCommandPayload(string, json.RawMessage) (any, error)
}
