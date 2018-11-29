package routes

import (
	"database/sql"
	"fmt"

	controller "../controllers"
	middleware "../middlewares"
	"github.com/gorilla/mux"
)

// API version and resource URI constants
const (
	version       = "v1"
	usersResource = "users"
	notesResource = "notes"
)

// Routes struct to initialize routes
type Routes struct{}

// Get takes in a Gorrilla mux router, controller, middleware and db handler to assemble
// an endpoint and its handler
func (routes *Routes) Get(r *mux.Router, c controller.Controller, m middleware.Middleware, db *sql.DB) {

	// OPEN ROUTES

	// URL: /v1/users/login
	r.Handle(fmt.Sprintf("/%v/%v/%v", version, usersResource, "authenticate"), c.LoginUser(db)).Methods("POST")

	// PROTECTED ROUTES

	// URL: /v1/notes
	r.Handle(fmt.Sprintf("/%v/%v", version, notesResource), m.Authorized(c.GetNotes(db))).Methods("GET")

	// URL: /v1/notes/:id
	r.Handle(fmt.Sprintf("/%v/%v/%v", version, notesResource, "{id}"), m.Authorized(c.GetNote(db))).Methods("GET")

	// URL: /v1/notes/:id
	r.Handle(fmt.Sprintf("/%v/%v", version, notesResource), m.Authorized(c.AddNote(db))).Methods("POST")

	// URL: /v1/notes/:id
	r.Handle(fmt.Sprintf("/%v/%v", version, notesResource), m.Authorized(c.UpdateNote(db))).Methods("PUT")

	// URL: /v1/notes/:id
	r.Handle(fmt.Sprintf("/%v/%v/%v", version, notesResource, "{id}"), m.Authorized(c.DeleteNote(db))).Methods("DELETE")
}
