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
		// note and slice of notes will need to be passed to the repo
		var note models.Note
		notes = []models.Note{}
		// Create repo
		notesRepo := notesrepository.NotesRepository{}

		// notes now have all notes from db
		notes, err := notesRepo.GetNotes(db, note, notes)
		if err != nil {
			// No notes returned
			h.JSONError(w, http.StatusNotFound, "Empty set")
			return
		}

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
			h.JSONError(w, http.StatusNotFound, "Note not found")
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
		noteID, err := noteRepo.AddNote(db, note)
		// No rows affected
		if err != nil {
			h.JSONError(w, http.StatusBadRequest, "Insert unsuccessful")
			return
		}

		// Respond with last inserted note ID
		h.JSONResponse(w, http.StatusOK, noteID)
	}
}

// UpdateNote updates a note and either returns the JSON of the
// affected row or an error.
func (c Controller) UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}

		rowsAffected, err := noteRepo.UpdateNote(db, note)
		// No rows were affected
		if err != nil {
			h.JSONError(w, http.StatusBadRequest, "Update unsuccessful")
			return
		}

		// Respond with number of rows affected
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
			h.JSONError(w, http.StatusBadRequest, "Delete unsuccessful")
			return
		}

		// Respond with number of rows affected
		h.JSONResponse(w, http.StatusOK, rowsAffected)
	}
}
