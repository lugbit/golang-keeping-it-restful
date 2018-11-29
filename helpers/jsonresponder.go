// Package helpers is a collection of common/helper functions
package helpers

import (
	"encoding/json"
	"net/http"

	"../models"
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

// JSONErrResponse
func JSONErrResponse(w http.ResponseWriter, resCode int, errorObjects []models.ErrorObject) {
	JSONResponse(w, resCode, models.ErrorsPayload{Errors: errorObjects})
}
