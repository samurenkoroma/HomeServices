package response

import (
	"encoding/json"
	"net/http"
)

func Json(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	response(writer, data, statusCode)
}
func ErrJson(writer http.ResponseWriter, msg string, statusCode int) {
	writer.Header().Set("Content-Type", "application/json")
	Json(writer, map[string]string{"error": msg}, statusCode)
}

func response(writer http.ResponseWriter, data any, statusCode int) {
	writer.Header().Set("Polevod-ID", "0001")
	writer.WriteHeader(statusCode)
	json.NewEncoder(writer).Encode(data)
}
