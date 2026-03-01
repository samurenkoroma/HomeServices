package dto

import "encoding/json"

type QueryPayload struct {
	Query string          `json:"query"`
	Data  json.RawMessage `json:"data"`
}
