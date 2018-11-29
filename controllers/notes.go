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
	validate "../validators"
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
		mapClaims, _ := h.ExtractClaims(r)
		// Extract the user ID from the JWT payload
		//
		// The value of mapClaims["id"] is type interface so we need
		// to do type assertion to float64 and then convert it to int.
		userID := int(mapClaims["id"].(float64))

		// note and slice of notes will need to be passed to the repo
		var note models.Note
		notes = []models.Note{}
		// Create repo
		notesRepo := notesrepository.NotesRepository{}

		// notes now have all notes from db
		notes, err := notesRepo.GetNotes(db, note, notes, userID)
		if err != nil {
			// No notes returned
			// New Error
			errObj := models.NewError(models.AppCodeNoteColEmpty, "No notes found", "Ensure you have at least one note", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
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
		mapClaims, _ := h.ExtractClaims(r)
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
			errObj := models.NewError(models.AppCodeNoteNotFound, "Note doesn't exist", "There is no note belonging to you with that ID", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
			return
		}

		h.JSONResponse(w, http.StatusOK, note)
	}
}

// AddNote adds a new note and then immediately calls GetNote to return
// the note to the client.
func (c Controller) AddNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := h.ExtractClaims(r)
		userID := int(mapClaims["id"].(float64))

		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		// Validate fields
		noteValidator := validate.Note{
			Title: note.Title,
			Note:  note.Note,
		}
		// Call the Validate method on the login validator
		// which returns a slice of ErrorObjects
		noteErrors := noteValidator.ValidateAddNote()
		// Check for errors
		if len(noteErrors) > 0 {
			// Respond with errors JSON
			h.JSONErrResponse(w, http.StatusBadRequest, noteErrors)
			return
		}

		noteRepo := notesrepository.NotesRepository{}
		// Pass note to AddNote to insert to DB
		lastInsertID, err := noteRepo.AddNote(db, note, userID)
		// No rows affected
		if err != nil {
			// Add unsuccessful
			errObj := models.NewError(models.AppCodeAddFailed, "Unable to add note", "Internal server error", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusInternalServerError, []models.ErrorObject{errObj})
			return
		}
		// Get inserted note to return to client
		note, err = noteRepo.GetNote(db, note, lastInsertID, userID)
		if err != nil {
			errObj := models.NewError(models.AppCodeNoteNotFound, "Note doesn't exist", "There is no note belonging to you with that ID", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
			return
		}

		// Respond with last inserted note
		h.JSONResponse(w, http.StatusCreated, note)
	}
}

// UpdateNote updates a user's existing note and responds with the updated note.
func (c Controller) UpdateNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := h.ExtractClaims(r)
		userID := int(mapClaims["id"].(float64))

		var note models.Note

		// Decode JSON body from request and store in note
		json.NewDecoder(r.Body).Decode(&note)

		// Validate fields
		noteValidator := validate.Note{
			ID:    note.ID,
			Title: note.Title,
			Note:  note.Note,
		}
		noteErrors := noteValidator.ValidateUpdateNote()
		if len(noteErrors) > 0 {
			h.JSONErrResponse(w, http.StatusBadRequest, noteErrors)
			return
		}

		noteRepo := notesrepository.NotesRepository{}

		// Attempts to update note, throw the number of rows affected returned
		_, err := noteRepo.UpdateNote(db, note, userID)
		if err != nil {
			// Update unsuccessful
			errObj := models.NewError(models.AppCodeNoteNotFound, "Updating note failed", "Ensure the note ID you are updating exists.", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
			return
		}

		// Updating note successful, return note.
		noteUpdated, err := noteRepo.GetNote(db, note, note.ID, userID)
		if err != nil {
			errObj := models.NewError(models.AppCodeNoteNotFound, "Note doesn't exist", "There is no note belonging to you with that ID", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
			return
		}
		// Respond with last updated note
		h.JSONResponse(w, http.StatusCreated, noteUpdated)
	}
}

// DeleteNote deletes a note and either returns the JSON of the
// affected row or an error.
func (c Controller) DeleteNote(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the JWT.
		mapClaims, _ := h.ExtractClaims(r)
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
			errObj := models.NewError(models.AppCodeNoteNotFound, "Deleting note failed", "Ensure the ID of the note you are trying to delete exists.", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusNotFound, []models.ErrorObject{errObj})
			return
		}

		// Respond with status no content (rows affected doesn't get shown)
		h.JSONResponse(w, http.StatusNoContent, rowsAffected)
	}
}
