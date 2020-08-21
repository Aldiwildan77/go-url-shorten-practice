package utils

import (
	"encoding/json"
	"net/http"
)

// BaseController is the struct for controller
type Responses struct{}

// ResponseWithError is the method for error message
func (base *Responses) ResponseWithError(w http.ResponseWriter, code int, message string) {
	base.ResponseWithJSON(w, code, map[string]string{"error": message})
}

// ResponseWithJSON is the method for response error using json
func (base *Responses) ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
