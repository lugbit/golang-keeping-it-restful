package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"./controllers"
	"./drivers"
	"./middlewares"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	gotenv "github.com/subosito/gotenv"
)

// Gobal DB handler variable
var db *sql.DB

func init() {
	// Load environment variables on startup
	err := gotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// Open DB handler and defer close
	db = drivers.ConnectDB()
	defer db.Close()

	// Initialize controllers
	controller := controllers.Controller{}
	// Initialize middlewares
	middleware := middlewares.Middleware{}

	// Gorilla mux
	r := mux.NewRouter()

	// Routes

	// Open routes
	r.Handle("/users/login", controller.LoginUser(db)).Methods("POST")

	// Protected routes
	r.Handle("/notes", middleware.Authorized(controller.GetNotes(db))).Methods("GET")
	r.Handle("/notes/{id}", middleware.Authorized(controller.GetNote(db))).Methods("GET")
	r.Handle("/notes", middleware.Authorized(controller.AddNote(db))).Methods("POST")
	r.Handle("/notes", middleware.Authorized(controller.UpdateNote(db))).Methods("PUT")
	r.Handle("/notes/{id}", middleware.Authorized(controller.DeleteNote(db))).Methods("DELETE")

	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))
}
