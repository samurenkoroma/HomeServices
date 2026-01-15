package request

import (
	"encoding/json"
	"io"
)

func decode[T any](body io.Reader) (T, error) {
	var payload T

	if err := json.NewDecoder(body).Decode(&payload); err != nil {
		return payload, err
	}
	return payload, nil
}
