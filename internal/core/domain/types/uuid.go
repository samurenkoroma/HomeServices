package types

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.New().String()
}

func UUIDIsValid(u string) bool {
	return uuid.Validate(u) == nil
}
