package validators

import (
	"strings"

	"../models"
)

// Note struct for initialization
type Note struct {
	// ID not required for Add note validation
	ID    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Note  string `json:"note"`
	// Colour not required for Add note validation
	Colour string               `json:"colour,omitempty"`
	Errors []models.ErrorObject `json:"errors"`
}

// ValidateAddNote validates the incoming JSON data coming from the client and returns
// a slice of ErrorObjects
func (v *Note) ValidateAddNote() []models.ErrorObject {
	v.Errors = []models.ErrorObject{}

	if strings.TrimSpace(v.Title) == "" {
		// Create new error object
		ErrObj := models.NewError(models.AppCodeNoteTitleEmpty, "Title cannot be empty", "Note title is required", "https://api.lugbit.com/docs/errors")
		// Append error to errors
		v.Errors = append(v.Errors, ErrObj)
	}

	if strings.TrimSpace(v.Note) == "" {
		ErrObj := models.NewError(models.AppCodeNoteTextEmpty, "Note text cannot be empty", "Note body is required", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	// Return slice of errors
	return v.Errors
}

// ValidateUpdateNote validates the incoming JSON data coming from the client and returns
// a slice of ErrorObjects
func (v *Note) ValidateUpdateNote() []models.ErrorObject {
	v.Errors = []models.ErrorObject{}

	// Updating a note requires a note ID greater than zero
	if v.ID <= 0 {
		ErrObj := models.NewError(models.AppCodeNoteIDEmpty, "Note ID is required", "A note ID is required when updating a note", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	if strings.TrimSpace(v.Title) == "" {
		ErrObj := models.NewError(models.AppCodeNoteTitleEmpty, "Title cannot be empty", "Note title is required", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	if strings.TrimSpace(v.Note) == "" {
		ErrObj := models.NewError(models.AppCodeNoteTextEmpty, "Note text cannot be empty", "Note body is required", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	return v.Errors
}
