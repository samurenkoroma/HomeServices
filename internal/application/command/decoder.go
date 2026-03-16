package command

import "encoding/json"

func DecodeCmd[T any](data []byte) (any, error) {
	var cmd T
	if err := json.Unmarshal(data, &cmd); err != nil {
		return nil, err
	}
	return cmd, nil
}
