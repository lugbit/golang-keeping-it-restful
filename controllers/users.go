// Package controllers stores handlers for routes
package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	h "../helpers"
	models "../models"
	"../repo/usersrepository"
	jwt "github.com/dgrijalva/jwt-go"
)

// LoginUser authenticates a user and generates and sends a signed JWT.
func (c Controller) LoginUser(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// user will hold the decoded JSON from the body
		user := models.User{}
		// Decode JSON from request body into user
		json.NewDecoder(r.Body).Decode(&user)

		// Init new users repo
		usersRepo := usersrepository.UsersRepository{}

		// Check if the provided email exist in the database
		//
		// userDB will return the user that the email belongs to and an error
		userDB, err := usersRepo.UserEmail(db, user.Email, user)
		if err != nil {
			// Email doesn't exist
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}

		// User with that email exists

		// Check if the password provided by the user matches that of
		// the password associated with the user returned by UserEmail.
		err = usersRepo.VerifyPassword(db, user.Password, userDB)
		if err != nil {
			// Password don't match
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		// User authenticated successfully, create JWT.

		// Create new claims with the user details
		userClaim := models.UserClaims{
			userDB.ID,
			userDB.FirstName,
			userDB.LastName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
				Issuer:    "Marck's RESTful API",
			},
		}
		// Generate signed JWT with claims
		signedJWT := h.GenerateJWT(userClaim)
		// Send signed token
		json.NewEncoder(w).Encode(signedJWT)
	})
}
