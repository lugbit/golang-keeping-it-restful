// Package helpers is a collection of common/helper functions
package helpers

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sets up response http headers and responds with JSON.
func JSONResponse(w http.ResponseWriter, resCode int, data interface{}) {
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	// Set http header to argument status code
	w.WriteHeader(resCode)
	// Encode data to JSON and write to response
	json.NewEncoder(w).Encode(data)
}

// JSONErrResponse piggy backs off JSONResponse and instead accepts
// only a message string.
func JSONErrResponse(w http.ResponseWriter, resCode int, msg string) {
	JSONResponse(w, resCode, map[string]string{"message:": msg})
}
