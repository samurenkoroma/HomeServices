package command

import "context"

type Command interface{}

type CommandHandler interface {
	Handle(ctx context.Context, cmd any) error
}
type CommandDecoder func([]byte) (any, error)
