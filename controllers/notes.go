// Package controllers stores handlers for routes
package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"../models"
	"../repo/notesrepository"
	"github.com/gorilla/mux"
)

// Slice of notes
var notes []models.Note

// Controller struct for initializing
type Controller struct{}

// GetNotes gets all notes
func (c Controller) GetNotes(db *sql.DB) http.HandlerFunc {
	// Return handlerfunc
	return func(w http.ResponseWriter, r *http.Request) {
		// note and slice of notes will need to be passed to the repo
		var note models.Note
		notes = []models.Note{}
		// Create repo
		notesRepo := notesrepository.NotesRepository{}

		// notes now have all notes from db
		notes = notesRepo.GetNotes(db, note, notes)

		// Encode slice of note to JSON and send as response
		json.NewEncoder(w).Encode(notes)
	}
}

// GetNote retrieves a singular note based on ID
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
		note = notesRepo.GetNote(db, note, id)

		json.NewEncoder(w).Encode(note)
	}
}

// AddNote adds a new note
func (c Controller) AddNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}
		// Pass note to AddNote to insert to DB
		noteID := noteRepo.AddNote(db, note)

		// Respond with last inserted note ID
		json.NewEncoder(w).Encode(noteID)
	}
}

// UpdateNote updates an existing note
func (c Controller) UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		noteRepo := notesrepository.NotesRepository{}
		rowsAffected := noteRepo.UpdateNote(db, note)

		// Respond with number of rows affected
		json.NewEncoder(w).Encode(rowsAffected)
	}
}

// DeleteNote deletes an existing note
func (c Controller) DeleteNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		// Convert id param to int
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			log.Fatalln(err)
		}

		noteRepo := notesrepository.NotesRepository{}
		rowsAffected := noteRepo.DeleteNote(db, id)

		// Respond with number of rows affected
		json.NewEncoder(w).Encode(rowsAffected)
	}
}
