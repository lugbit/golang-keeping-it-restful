package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Middleware is a struct used to initialize and call middleware functions
type Middleware struct{}

// Authorized is a middleware that checks request headers for a valid JWT token
func (m Middleware) Authorized(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retrieve secret environment variable
		hmacSecret := []byte(os.Getenv("JWT_SECRET"))

		// Check request header for a Token
		// Check if the length of the Authorization value is greater than 0
		if len(r.Header.Get("Authorization")) > 0 {

			// Get JWT from the Authorization header. As it gives us
			// bearer JWT, we need to split the JWT.
			authSlice := strings.Split(r.Header.Get("Authorization"), " ")
			jwtToken := authSlice[1]

			// Parse token in header
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				// Check that the algorithm used in parsed token is the same as what was issued by the server
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				// Return the secret and nil error
				return hmacSecret, nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// If token is valid, server the argument handler
			if token.Valid {
				h.ServeHTTP(w, r)
			}
		} else {
			// No token in header
			http.Error(w, "Not Authorized", http.StatusForbidden)
			return
		}

	})
}
