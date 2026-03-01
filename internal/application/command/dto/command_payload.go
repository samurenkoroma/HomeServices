package dto

import "encoding/json"

type CommandPayload struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}
