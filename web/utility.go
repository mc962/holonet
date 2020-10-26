package web

import (
	"encoding/json"
	"net/http"
)

func writeJSON(writer http.ResponseWriter, response interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Accept", "application/json")

	_ = json.NewEncoder(writer).Encode(response)
}
