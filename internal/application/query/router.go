package query

import "context"

type QueryRouter interface {
	Dispatch(ctx context.Context, queryName string, query any) (any, error)
	Decode(queryName string, data []byte) (any, error)
}
