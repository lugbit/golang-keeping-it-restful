// Package controllers stores handlers for routes
package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	h "../helpers"
	"../models"
	"../repo/notesrepository"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// Slice of notes
var notes []models.Note

// Controller is a struct for initializing and calling controllers
type Controller struct{}

// GetNotes gets every note belonging to a user and either returns the JSON of the rows
// or an error.
func (c Controller) GetNotes(db *sql.DB) http.HandlerFunc {
	// Return handlerfunc
	return func(w http.ResponseWriter, r *http.Request) {

		// Extract claims from JWT
		mapClaims, _ := extractClaims(r)
		// Extract the user ID from the JWT payload
		//
		// The value of mapClaims["id"] is type interface so we need
		// to do type assertion to float64 and then convert it to int.
		userID := int(mapClaims["id"].(float64))

		//note and slice of notes will need to be passed to the repo
		var note models.Note
		notes = []models.Note{}
		// Create repo
		notesRepo := notesrepository.NotesRepository{}

		// notes now have all notes from db
		notes, err := notesRepo.GetNotes(db, note, notes, userID)
		if err != nil {
			// No notes returned
			h.JSONErrResponse(w, http.StatusNotFound, "empty collection")
			return
		}

		// Respond with status OK and collection of notes as payload
		h.JSONResponse(w, http.StatusOK, notes)
	}
}

// GetNote gets a particular belonging to a user and either returns
// the JSON of the row or an error.
func (c Controller) GetNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := extractClaims(r)
		userID := int(mapClaims["id"].(float64))

		var note models.Note
		// Collect variable(s) from the URL
		params := mux.Vars(r)

		notesRepo := notesrepository.NotesRepository{}

		// Convert id parameter from string to int
		noteID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}
		// GetNote returns a singular models.Note
		note, err = notesRepo.GetNote(db, note, noteID, userID)
		if err != nil {
			// No note returned
			h.JSONErrResponse(w, http.StatusNotFound, "resource not found")
			return
		}

		h.JSONResponse(w, http.StatusOK, note)
	}
}

// AddNote adds a new note and either returns the JSON last insert
// ID or an error.
func (c Controller) AddNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := extractClaims(r)
		userID := int(mapClaims["id"].(float64))

		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}
		// Pass note to AddNote to insert to DB
		lastInsertID, err := noteRepo.AddNote(db, note, userID)
		// No rows affected
		if err != nil {
			h.JSONErrResponse(w, http.StatusBadRequest, "add unsuccessful")
			return
		}

		// Respond with last inserted note ID
		h.JSONResponse(w, http.StatusCreated, lastInsertID)
	}
}

// UpdateNote updates a user's existing note OR creates a new note if the
// note does not exist.
func (c Controller) UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := extractClaims(r)
		userID := int(mapClaims["id"].(float64))

		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}

		// Attempts to update note
		rowsAffected, err := noteRepo.UpdateNote(db, note, userID)
		// Affected rows is zero, no notes were updated,
		// create resource instead
		if err != nil {
			lastInsertID, err := noteRepo.AddNote(db, note, userID)
			if err != nil {
				// Insert failed, respond with error
				h.JSONErrResponse(w, http.StatusBadRequest, "add unsuccessful")
				return
			}
			// Update failed but insert is successful. Respond with
			// status code created and last insert ID.
			h.JSONResponse(w, http.StatusCreated, lastInsertID)
			return
		}

		// Updating existing note successful, respond with number
		// of rows affected.
		h.JSONResponse(w, http.StatusOK, rowsAffected)
	}
}

// DeleteNote deletes a note and either returns the JSON of the
// affected row or an error.
func (c Controller) DeleteNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := extractClaims(r)
		userID := int(mapClaims["id"].(float64))

		params := mux.Vars(r)

		// Convert id param to int
		noteID, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}

		noteRepo := notesrepository.NotesRepository{}
		rowsAffected, err := noteRepo.DeleteNote(db, noteID, userID)
		if err != nil {
			h.JSONErrResponse(w, http.StatusBadRequest, "delete unsuccessful")
			return
		}

		// Respond with status code no content
		h.JSONResponse(w, http.StatusNoContent, rowsAffected)
	}
}

func extractClaims(r *http.Request) (jwt.MapClaims, bool) {
	// Retrieve secret environment variable
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))

	// Split JWT from the request authorization header
	authSlice := strings.Split(r.Header.Get("Authorization"), " ")
	// JWT will be at index 1 after the "Bearer"
	jwtToken := authSlice[1]

	token, err := jwt.Parse(jwtToken, func(jwtToken *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}
	// Check if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	// Invalid JWT
	return nil, false
}
