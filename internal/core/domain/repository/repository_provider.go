package repository

import "errors"

type RepositoryProvider interface {
	ProviderName() string
}

var (
	ErrInvalidProviderType = errors.New("invalid provider type")
)
