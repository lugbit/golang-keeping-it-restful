// Package controllers stores handlers for routes
package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	h "../helpers"
	"../models"
	"../repo/notesrepository"
	"github.com/gorilla/mux"
)

// Slice of notes
var notes []models.Note

// Controller struct for initializing
type Controller struct{}

// GetNotes gets all notes and either returns the JSON of the rows
// or an error.
func (c Controller) GetNotes(db *sql.DB) http.HandlerFunc {
	// Return handlerfunc
	return func(w http.ResponseWriter, r *http.Request) {
		//note and slice of notes will need to be passed to the repo
		var note models.Note
		notes = []models.Note{}
		// Create repo
		notesRepo := notesrepository.NotesRepository{}

		// notes now have all notes from db
		notes, err := notesRepo.GetNotes(db, note, notes)
		if err != nil {
			// No notes returned
			h.JSONErrResponse(w, http.StatusNotFound, "empty collection")
			return
		}

		// Respond with status OK and collection of notes as payload
		h.JSONResponse(w, http.StatusOK, notes)
	}
}

// GetNote gets a particular note based on the ID and either returns
// the JSON of the row or an error.
func (c Controller) GetNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note
		// Collect variable(s) from the URL
		params := mux.Vars(r)

		notesRepo := notesrepository.NotesRepository{}

		// Convert id parameter from string to int
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}
		// GetNote returns a singular models.Note
		note, err = notesRepo.GetNote(db, note, id)
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
		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}
		// Pass note to AddNote to insert to DB
		lastInsertID, err := noteRepo.AddNote(db, note)
		// No rows affected
		if err != nil {
			h.JSONErrResponse(w, http.StatusBadRequest, "add unsuccessful")
			return
		}

		// Respond with last inserted note ID
		h.JSONResponse(w, http.StatusCreated, lastInsertID)
	}
}

// UpdateNote updates an existing note OR creates a new note if the
// resource doesn't exist.
func (c Controller) UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}

		// Attempts to update note
		rowsAffected, err := noteRepo.UpdateNote(db, note)
		// Affected rows is zero, no notes were updated,
		// create resource instead
		if err != nil {
			lastInsertID, err := noteRepo.AddNote(db, note)
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
		params := mux.Vars(r)

		// Convert id param to int
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}

		noteRepo := notesrepository.NotesRepository{}
		rowsAffected, err := noteRepo.DeleteNote(db, id)
		if err != nil {
			h.JSONErrResponse(w, http.StatusBadRequest, "delete unsuccessful")
			return
		}

		// Respond with status code no content
		h.JSONResponse(w, http.StatusNoContent, rowsAffected)
	}
}
