// Package validators is for validating incoming JSON data from the client such as when authenticating or adding a new note.
package validators

import (
	"regexp"
	"strings"

	"../models"
)

// Login struct for initialization
type Login struct {
	Email    string               `json:"email"`
	Password string               `json:"password"`
	Errors   []models.ErrorObject `json:"errors"`
}

// Validate validates the incoming JSON data coming from the client and returns
// a slice of ErrorObjects
func (v *Login) Validate() []models.ErrorObject {
	v.Errors = []models.ErrorObject{}

	// Check email is not empty and if the correct format
	if strings.TrimSpace(v.Email) != "" {
		// Check if the correct format using regexp
		re := regexp.MustCompile(".+@.+\\..+")
		matched := re.Match([]byte(v.Email))

		if matched == false {
			ErrObj := models.NewError(models.AppCodeEmailInvalid, "Email is invalid", "Ensure the email is formatted in example@email.com", "https://api.lugbit.com/docs/errors")
			v.Errors = append(v.Errors, ErrObj)
		}
	} else {
		// Email is empty
		ErrObj := models.NewError(models.AppCodeEmailEmpty, "Email cannot be empty", "Email is a required field", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	// Check if password is empty
	if strings.TrimSpace(v.Password) == "" {
		ErrObj := models.NewError(models.AppCodePasswordEmpty, "Password cannot be empty", "Password is a required field", "https://api.lugbit.com/docs/errors")
		v.Errors = append(v.Errors, ErrObj)
	}

	// Return slice of errors
	return v.Errors
}
