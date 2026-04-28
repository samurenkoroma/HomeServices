package command

import (
	"context"
	"encoding/json"
)

type Handler func(ctx context.Context, cmd any) (any, error)
type DecoderFunc func([]byte) (any, error)

type Router interface {
	Register(string, Handler, DecoderFunc)
	Dispatch(ctx context.Context, commandName string, cmd any) (any, error)
	ResolveCommandPayload(string, json.RawMessage) (any, error)
}
