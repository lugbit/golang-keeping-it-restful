package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"./controllers"
	"./drivers"
	mw "./middlewares"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	gotenv "github.com/subosito/gotenv"
)

// Gobal DB handler variable
var db *sql.DB

func init() {
	// Load environment variables on startup
	gotenv.Load()
}

func main() {
	// Open DB handler and defer close
	db = drivers.ConnectDB()
	defer db.Close()

	// New controller
	controller := controllers.Controller{}

	// Gorilla mux
	r := mux.NewRouter()

	// Routes
	r.Handle("/notes", mw.VerifyToken(controller.GetNotes(db))).Methods("GET")
	r.HandleFunc("/notes/{id}", controller.GetNote(db)).Methods("GET")
	r.HandleFunc("/notes", controller.AddNote(db)).Methods("POST")
	r.HandleFunc("/notes", controller.UpdateNote(db)).Methods("PUT")
	r.HandleFunc("/notes/{id}", controller.DeleteNote(db)).Methods("DELETE")
	r.Handle("/users/login", controller.LoginUser(db)).Methods("POST")

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))
}
