// Package helpers is a collection of common/helper functions
package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse responds with JSON after setting up appropriate headers
func JSONResponse(w http.ResponseWriter, resCode int, data interface{}) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	w.Write(response)
}

// JSONError responds with an error message
func JSONError(w http.ResponseWriter, resCode int, msg string) {
	JSONResponse(w, resCode, map[string]string{"message": msg})
}
