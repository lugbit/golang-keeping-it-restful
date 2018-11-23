package helpers

import (
	"log"
	"os"

	models "../models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT takes in a claims then generates and signs a JWT.
func GenerateJWT(claims models.UserClaims) string {
	// Generate token with HS256 algorithm. Note: this is not yet signed and
	// only the header and payload will be base64 encoded.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with the secret key. This will hash the header . payload with the secret
	// to create the signature.
	ss, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Fatalln(err)
	}
	// Return signed token string
	return ss
}
