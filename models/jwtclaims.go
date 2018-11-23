package models

import jwt "github.com/dgrijalva/jwt-go"

// UserClaims is a custom JWT claims struct which is used
// by the jwt-go package.
//
// This is the 'payload' part of the generated JWT sent to
// the user after successfully authenticating.
type UserClaims struct {
	// Custom claims
	UserID    int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Standard claims (part of jwt-go)
	jwt.StandardClaims
}
