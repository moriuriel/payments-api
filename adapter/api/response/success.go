package response

import (
	"encoding/json"
	"net/http"
	"time"
)

type (
	Success struct {
		statusCode int
		result     interface{}
	}
	Response struct {
		Content    interface{} `json:"content"`
		StatusCode int         `json:"statusCode"`
		Timestamp  string      `json:"timestamp"`
	}
)

func NewSuccess(result interface{}, status int) Success {
	return Success{
		statusCode: status,
		result:     result,
	}
}

func (r Success) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.statusCode)

	response := Response{
		StatusCode: r.statusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
		Content:    r.result,
	}

	return json.NewEncoder(w).Encode(response)
}
