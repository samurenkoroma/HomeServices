package utils

import "encoding/json"

func DecodeJSON[T any](data []byte) (any, error) {
	var v T
	err := json.Unmarshal(data, &v)
	return &v, err
}
