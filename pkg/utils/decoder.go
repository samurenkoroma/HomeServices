package utils

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func DecodeJSON[T any](data []byte) (any, error) {
	var v T
	validate := validator.New()
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	return &v, validate.Struct(v)
}
