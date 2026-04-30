package query

import (
	"context"
	"errors"
)

// Decoder преобразует сырой payload (json, grpc, etc)
// в конкретный query struct
type Decoder func([]byte) (any, error)
type Handler func(ctx context.Context, payload any) (any, error)

//type HandlerFunc func(ctx context.Context, payload any) (any, error)

type Router interface {
	// Register регистрирует query
	Register(string, Handler, Decoder)
	// Dispatch выполняет query
	Dispatch(ctx context.Context, name string, payload []byte) (any, error)
}

var ErrQueryNotFound = errors.New("query not registered")
