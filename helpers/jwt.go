package helpers

import (
	"log"
	"net/http"
	"os"
	"strings"

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

// ExtractClaims extracts the JWT claims and returns the map and error
func ExtractClaims(r *http.Request) (jwt.MapClaims, bool) {
	// Retrieve secret environment variable
	hmacSecret := []byte(os.Getenv("JWT_SECRET"))

	// Split JWT from the request authorization header
	authSlice := strings.Split(r.Header.Get("Authorization"), " ")
	// JWT will be at index 1 after the "Bearer"
	jwtToken := authSlice[1]

	token, err := jwt.Parse(jwtToken, func(jwtToken *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}
	// Check if token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	// Invalid JWT
	return nil, false
}
