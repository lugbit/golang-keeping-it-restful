// Package controllers stores handlers for routes
package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	h "../helpers"
	"../models"
	"../repo/usersrepository"
	validate "../validators"
	jwt "github.com/dgrijalva/jwt-go"
)

// LoginUser authenticates a user and generates and sends a signed JWT.
func (c Controller) LoginUser(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user will hold the decoded JSON from the body
		user := models.User{}
		// Decode JSON from request body into user
		json.NewDecoder(r.Body).Decode(&user)

		// Validate fields
		login := validate.Login{
			Email:    user.Email,
			Password: user.Password,
		}
		// Call the Validate method on the login validator
		// which returns a slice of ErrorObjects
		loginErrors := login.Validate()
		// Check for errors
		if len(loginErrors) > 0 {
			// Respond with errors JSON
			h.JSONErrResponse(w, http.StatusBadRequest, loginErrors)
			return
		}

		// Validation passed

		// Init new users repo
		usersRepo := usersrepository.UsersRepository{}

		// Check if the provided email exist in the database
		//
		// userDB will return the user that the email belongs to and an error
		userDB, err := usersRepo.UserEmail(db, user.Email, user)
		if err != nil {
			errObj := models.NewError(models.AppCodeEmailNotFound, "Incorrect username or password", "Verify that the email address and password are correct", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusBadRequest, []models.ErrorObject{errObj})
			return
		}

		// User with that email exists

		// Check if the password provided by the user matches that of
		// the password associated with the user returned by UserEmail.
		err = usersRepo.VerifyPassword(db, user.Password, userDB)
		if err != nil {
			// Password don't match
			errObj := models.NewError(models.AppCodePasswordIncorrect, "Incorrect username or password", "Verify that the email address and password are correct", "https://api.lugbit.com/docs/errors")
			h.JSONErrResponse(w, http.StatusBadRequest, []models.ErrorObject{errObj})
			return
		}

		// User authenticated successfully, create JWT.

		// Create new claims with the user details
		userClaim := models.UserClaims{
			UserID:    userDB.ID,
			FirstName: userDB.FirstName,
			LastName:  userDB.LastName,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Issuer:    "Marck's RESTful API",
			},
		}
		// Generate signed JWT with claims
		signedJWT := h.GenerateJWT(userClaim)
		// Send signed token
		h.JSONResponse(w, http.StatusOK, models.JWT{Token: signedJWT})
	})
}
